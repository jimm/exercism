defmodule Palindromes do

  @doc """
  Generates all palindrome products from an optionally given min factor (or 1) to a given max factor.
  """
  @spec generate(non_neg_integer, non_neg_integer) :: map() 
  def generate(max_factor, min_factor \\ 1) do
    pals = %{}
    for a <- (min_factor..max_factor),
        b <- (min_factor..max_factor),
        prod = a*b,
        palindrome?(prod) do
      pals = Map.put(pals, prod, [[a,b] | Map.get(pals, prod, [])])
      IO.puts "pals is now #{inspect pals}"
    end
    pals
  end

  def palindrome?(n) do
    s = Integer.to_string(n)
    s == String.reverse(s)
  end
end
