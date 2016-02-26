defmodule DNA do
  @doc """
  Returns number of differences between two strands of DNA, known as the Hamming Distance.

  ## Examples

  iex> DNA.hamming_distance('AAGTCATA', 'TAGCGATC')
  4
  """
  @spec hamming_distance([char], [char]) :: non_neg_integer
  def hamming_distance(chars1, chars2) do
    hamming_distance(chars1, chars2, 0)
  end

  def hamming_distance([], [], n), do: n
  def hamming_distance([], _, _), do: nil
  def hamming_distance(_, [], _), do: nil
  def hamming_distance([h|t1], [h|t2], n), do: hamming_distance(t1, t2, n)
  def hamming_distance([_|t1], [_|t2], n), do: hamming_distance(t1, t2, n+1)
end
