defmodule ETL do
  @doc """
  Transform an index into an inverted index.

  ## Examples

  iex> ETL.transform(%{"a" => ["ABILITY", "AARDVARK"], "b" => ["BALLAST", "BEAUTY"]})
  %{"ability" => "a", "aardvark" => "a", "ballast" => "b", "beauty" =>"b"}
  """
  @spec transform(Map.t) :: map()
  def transform(input), do: transform(Enum.into(input, []), %{})

  defp transform([], acc), do: acc
  defp transform([{_, []}|rest], acc), do: transform(rest, acc)
  defp transform([{k, [h|t]}|rest], acc) do
    transform([{k, t}|rest], Map.put(acc, String.downcase(h), k))
  end
end
