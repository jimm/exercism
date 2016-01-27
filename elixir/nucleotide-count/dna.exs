defmodule DNA do
  @nucleotides [?A, ?C, ?G, ?T]

  @doc """
  Counts individual nucleotides in a DNA strand.

  ## Examples

  iex> DNA.count('AATAA', ?A)
  4

  iex> DNA.count('AATAA', ?T)
  1
  """
  @spec count([char], char) :: non_neg_integer
  def count(strand, nucleotide) when nucleotide in @nucleotides do
    strand
    |> Enum.filter(&(legit_strand_char(&1) == nucleotide))
    |> length
  end
  def count(_, _) do
    raise ArgumentError
  end


  defp legit_strand_char(c) when c in @nucleotides, do: c
  defp legit_strand_char(_), do: raise ArgumentError

  @doc """
  Returns a summary of counts by nucleotide.

  ## Examples

  iex> DNA.histogram('AATAA')
  %{?A => 4, ?T => 1, ?C => 0, ?G => 0}
  """
  @spec histogram([char]) :: Dict.t
  def histogram(strand) do
    counts = @nucleotides
    |> Enum.map(&(count(strand, &1)))
    Enum.into(Enum.zip(@nucleotides, counts), %{})
  end
end
