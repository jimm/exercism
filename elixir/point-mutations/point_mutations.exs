defmodule DNA do
  @doc """

  Returns number of differences between two strands of DNA, known as the
  Hamming Distance.

  Well, this isn't really Hamming Distance. That would include insertions
  and deletions.

  ## Examples

  iex> DNA.hamming_distance('AAGTCATA', 'TAGCGATC')
  4
  """
  @spec hamming_distance([char], [char]) :: non_neg_integer
  def hamming_distance(strand1, strand2) when length(strand1) == length(strand2) do
    Enum.zip(strand1, strand2)
    |> Enum.filter(fn({a, b}) -> a != b end)
    |> length
  end
  def hamming_distance(_, _), do: nil
end
