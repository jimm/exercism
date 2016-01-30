defmodule Palindromes do

  @doc """
  Generates all palindrome products from an optionally given min factor (or 1) to a given max factor.
  """
  @spec generate(non_neg_integer, non_neg_integer) :: map() 
  def generate(max_factor, min_factor \\ 1) do
    pals = for a <- (min_factor..max_factor),
               b <- (min_factor..max_factor),
               prod = a*b,
               palindrome?(prod) do
             {prod, [a, b]}
           end
    condense_map(pals, %{})
  end

  defp condense_map([], m), do: m
  defp condense_map([{n, pair}|t], m) do
    IO.puts "n #{n}, pair #{inspect pair}, t #{inspect t}, m #{inspect m}"
    condense_map(t, Map.put(m, n, [Enum.sort(pair) | Map.get(m, n, [])]))
  end

  def palindrome?(n) do
    s = Integer.to_string(n)
    s == String.reverse(s)
  end
end
