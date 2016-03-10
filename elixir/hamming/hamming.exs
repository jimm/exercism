defmodule DNA do
  @doc """
  Returns number of differences between two strands of DNA, known as the Hamming Distance.

  ## Examples

  iex> DNA.hamming_distance('AAGTCATA', 'TAGCGATC')
  4
  """
  @spec hamming_distance([char], [char]) :: non_neg_integer
  def hamming_distance(chars1, chars2) when length(chars1) == length(chars2) do
    length(chars1 -- chars2)
  end
  def hamming_distance(_, _), do: nil
end
