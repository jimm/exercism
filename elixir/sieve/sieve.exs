defmodule Sieve do

  @doc """
  Generates a list of primes up to a given limit.
  """
  @spec primes_to(non_neg_integer) :: [non_neg_integer]
  def primes_to(limit) do
    nums = (2..limit) |> Enum.to_list
    primes_in(nums, [])
  end

  defp primes_in([], primes), do: Enum.reverse(primes)
  defp primes_in([h|t], primes) do
    t
    |> Enum.filter(fn(n) -> rem(n, h) != 0 end)
    |> primes_in([h|primes])
  end
end
