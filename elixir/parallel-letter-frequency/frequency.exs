defmodule WadingPool do
  @moduledoc """
  A very simple worker pool that uses Tasks.
  """

  @doc """
  Returns a pool of workers. 
  """
  def start(num_workers) do
    {:ok, {[], [], num_workers}}
  end

  # If there is a free slot, spawn a Task. If there is not, await the
  # earliest-created task. This isn't necessarily great since a task created
  # later may have already finished, but it's good enough for a wading pool.
  def run({running, results, max_workers}, f)
  when length(running) < max_workers
  do
    # We keep the oldest task at the start of the list by putting new tasks
    # at the end of the list.
    {running ++ [Task.async(f)], results, max_workers}
  end
  def run({[task|remaining], results, max_workers}, f) do
    run({remaining, results ++ [Task.await(task)], max_workers}, f)
  end

  def join({running, results, _}) do
    finished_results = running |> Enum.map(&Task.await/1)
    results ++ finished_results
  end
end

defmodule Frequency do
  @doc """
  Count word frequency in parallel.

  Returns a dict of characters to frequencies.

  The number of worker processes to use can be set with 'workers'.
  """
  @spec frequency([String.t], pos_integer) :: Dict.t
  def frequency(texts, workers) do
    {:ok, pool} = WadingPool.start(workers)
    pool = texts
    |> Enum.reduce(pool,
                   fn(str, pool) -> WadingPool.run(pool,
                                                   fn() -> frequencies(str) end)
      end)

    pool
    |> WadingPool.join
    |> Enum.reduce(%{}, fn(freqs, acc) ->
      add_frequencies(Enum.to_list(freqs), acc)
    end)
  end

  def frequencies(str) do
    chars = str
    |> String.replace(~r/[^[:alpha:]]/u, "")
    |> String.downcase
    |> String.graphemes
    frequencies(chars, %{})
  end

  defp frequencies([], freqs), do: freqs
  defp frequencies([h|t], freqs) do
    frequencies(t, Map.update(freqs, String.downcase(h), 1, &(&1 + 1)))
  end

  defp add_frequencies(freqs_list, results) do
    Enum.reduce(freqs_list, results, fn({str, n}, acc) ->
      Map.update(acc, str, n, &(&1 + n))
    end)
  end
end
