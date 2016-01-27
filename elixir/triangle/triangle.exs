defmodule Triangle do
  @type kind :: :equilateral | :isosceles | :scalene

  @doc """
  Return the kind of triangle of a triangle with 'a', 'b' and 'c' as lengths.
  """
  @spec kind(number, number, number) :: { :ok, kind } | { :error, String.t }
  def kind(a, a, a), do: check(a, a, a, :equilateral)
  def kind(a, a, b), do: check(a, a, b, :isosceles)
  def kind(a, b, a), do: check(a, b, a, :isosceles)
  def kind(b, a, a), do: check(b, a, a, :isosceles)
  def kind(a, b, c), do: check(a, b, c, :scalene)

  defp check(a, b, c, type) do
    cond do
      a <= 0 || b <= 0 || c <= 0 ->
        {:error, "all side lengths must be positive"}
      check_inequality(a, b, c) == false ->
        {:error, "side lengths violate triangle inequality"}
      true ->
        {:ok, type}
    end
  end

  defp check_inequality(a, b, c) do
    a + b > c && a + c > b && b + c > a
  end
end
