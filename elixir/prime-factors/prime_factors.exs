defmodule PrimeFactors do
  @doc """
  Compute the prime factors for 'number'.

  The prime factors are prime numbers that when multiplied give the desired
  number.

  The prime factors of 'number' will be ordered lowest to highest. 
  """
  @spec factors_for(pos_integer) :: [pos_integer]
  def factors_for(number) do
    factors_for(number, 2, [])
  end

  defp factors_for(1, _, prime_factors), do: Enum.reverse(prime_factors)
  defp factors_for(number, prime, prime_factors) do
    if rem(number, prime) == 0 do
      factors_for(div(number, prime), prime, [prime|prime_factors])
    else
      factors_for(number, next_prime(prime), prime_factors)
    end
  end

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
