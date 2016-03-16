defmodule Matrix do

  @min_int -1_000_000
  @max_int  1_000_000

  @doc """
  Parses a string representation of a matrix
  to a list of rows
  """
  @spec rows(String.t()) :: [[integer]]
  def rows(str) do
    str
    |> String.split("\n")
    |> Enum.map(&row_to_ints/1)
  end

  @doc """
  Parses a string representation of a matrix
  to a list of columns
  """
  @spec columns(String.t()) :: [[integer]]
  def columns(str) do
    str
    |> rows
    |> rotate
  end

  @doc """
  Calculates all the saddle points from a string
  representation of a matrix
  """
  @spec saddle_points(String.t()) :: [{integer, integer}]
  def saddle_points(str) do
    rs = rows(str)
    row_max_indexes_with_index = rs
    |> max_indexes
    |> Stream.with_index

    col_min_indexes = rs
    |> rotate
    |> min_indexes

    row_max_indexes_with_index
    |> Enum.flat_map(fn({row_max_indexes, row_index}) ->
      row_max_indexes
      |> Enum.map(fn(row_max_index) ->
        if Enum.member?(Enum.at(col_min_indexes, row_max_index), row_index) do
          {row_index, col_index}
        else
          nil
        end
      end)
    end)
    |> Enum.reject(&is_nil/1)
  end

  defp row_to_ints(str) do
    str
    |> String.strip
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  defp rotate(rows) do
    (0..length(hd(rows))-1)
    |> Enum.map(&column(&1, rows))
  end

  defp column(c, rows) do
    rows
    |> Enum.map(&Enum.at(&1, c))
  end

  defp max_indexes(rows) do
    comparator = fn
      (prev, prev) -> :keep
      (prev, new) when new > prev -> :start
      (_, _) -> :ignore
    end
    index_finder(rows, @min_int, comparator, [])
  end

  defp min_indexes(rows) do
    comparator = fn
      (prev, prev) -> :keep
      (prev, new) when new < prev -> :start
      (_, _) -> :ignore
    end
    index_finder(rows, @max_int, comparator, [])
  end

  # Returns list of index lists, one for each row. The function f must take
  # two values, previous and new, and return one of three values: :start,
  # :keep, or :ignore.
  defp index_finder([], _, _, indexes), do: Enum.reverse indexes
  defp index_finder([row|rows], starting_prev, f, indexes) do
    reducer = fn(elem, {i, prev_val, ixs}) ->
      case f.(prev_val, elem) do
        :start -> {i+1, elem, [i]}
        :keep -> {i+1, elem, [i|ixs]}
        :ignore -> {i+1, prev_val, ixs}
      end
    end
    reduction = row
    |> Enum.reduce({0, starting_prev, []}, reducer)

    {_, _, row_indexes} = reduction
    index_finder(rows, starting_prev, f, [row_indexes|indexes])
  end
end
