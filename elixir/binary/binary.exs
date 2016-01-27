defmodule Binary do
  @doc """
  Convert a string containing a binary number to an integer.

  On errors returns 0.
  """
  @spec to_decimal(String.t) :: non_neg_integer
  def to_decimal(string), do: to_decimal(string, 0)

  def to_decimal("", n), do: n
  def to_decimal("1" <> str, n), do: to_decimal(str, n * 2 + 1)
  def to_decimal("0" <> str, n), do: to_decimal(str, n * 2)
  def to_decimal(<<_::size(8)>> <> str, n), do: 0
end
