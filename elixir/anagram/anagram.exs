defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t, [String.t]) :: [String.t]
  def match(base, candidates) do
    candidates
    |> Enum.filter(&(anagram?(&1, base)))
  end

  defp anagram?(s1, s2) do
    ds1 = String.downcase(s1)
    ds2 = String.downcase(s2)
    sorted_chars(ds1) == sorted_chars(ds2) && ds1 != ds2
  end

  defp sorted_chars(s) do
    s
    |> String.graphemes
    |> Enum.sort
  end
end
