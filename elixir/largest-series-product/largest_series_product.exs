defmodule Series do

  @doc """
  Splits up the given string of numbers into an array of integers.
  """
  @spec digits(String.t) :: [non_neg_integer]
  def digits(number_string) do
    number_string
    |> String.graphemes
    |> Enum.map(&String.to_integer/1)
  end

  @doc """
  Generates sublists of a given size from a given string of numbers.
  """
  @spec slices(String.t, non_neg_integer) :: [list(non_neg_integer)]
  def slices(number_string, size) do
    if String.length(number_string) < size, do: raise ArgumentError
    number_string
    |> digits
    |> Enum.chunk(size, 1)
  end

  @doc """

  Finds the largest product of a given number of consecutive numbers in a
  given string of numbers.
  """
  @spec largest_product(String.t, non_neg_integer) :: non_neg_integer
  def largest_product(_, 0), do: 1
  def largest_product("", 0), do: 1
  def largest_product("", _), do: raise ArgumentError
  def largest_product(number_string, size) do
    slices(number_string, size)
    |> Enum.map(fn(slice) -> Enum.reduce(slice, 1, &*/2) end)
    |> Enum.max
  end
end
