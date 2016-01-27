defmodule Words do

  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t) :: map()
  def count(sentence) do
    words = sentence
    |> String.downcase
    |> String.replace(~r{[^-\w]+}u, " ")
    |> String.strip
    |> String.split(~r{[\s_]+})

    word_freqs(%{}, words)
  end

  def word_freqs(freqs, []), do: freqs
  def word_freqs(freqs, [word|words]) do
    word_freqs(Map.update(freqs, word, 1, &(&1 + 1)), words)
  end
end
