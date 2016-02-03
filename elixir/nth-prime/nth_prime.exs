defmodule Prime do

  @doc """
  Generates the nth prime.
  """
  @spec nth(non_neg_integer) :: non_neg_integer
  def nth(count), do: nth(count, 2)

  def nth(count, _) when count < 1, do: raise ArgumentError
  def nth(1, p), do: p
  def nth(count, p), do: nth(count-1, next_prime(p))

  defp next_prime(2), do: 3
  defp next_prime(n) do
    Stream.iterate(n + 2, &(&1 + 2))
    |> Stream.drop_while(&(!prime?(&1)))
    |> Enum.take(1)
    |> hd
  end

  defp prime?(n) when is_integer(n) do
    ceiling = round(:math.sqrt(n))
    if n == ceiling * ceiling do
      false
    else
      prime_test(n, 2, ceiling, ceiling)
    end
  end

  defp prime_test(_, i, i, h), do: i == h
  defp prime_test(n, i, j, h) do
    if rem(n, i) == 0 do
      i == h
    else
      prime_test(n, i + 1, j, h)
    end
  end
end