defmodule SumOfMultiples do
  @doc """
  Adds up all numbers from 1 to a given end number that are multiples of the factors provided.

  The default factors are 3 and 5.
  """
  @spec to(non_neg_integer, [non_neg_integer]) :: non_neg_integer
  def to(limit, factors \\ [3, 5]), do: to(limit, factors, [])

  defp to(_, [], multiples), do: multiples |> Enum.uniq |> Enum.sum
  defp to(limit, [factor|t], multiples) do
    factors = factor
    |> Stream.iterate(&(&1 + factor))
    |> Stream.take_while(fn(n) -> n < limit end)
    |> Enum.to_list
    to(limit, t, Enum.concat(factors, multiples))
  end
end
