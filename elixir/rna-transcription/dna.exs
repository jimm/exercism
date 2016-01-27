defmodule DNA do
  @transcriptions %{?A => ?U, ?C => ?G, ?G => ?C, ?T => ?A}

  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> DNA.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    dna |> Enum.map(&(Map.get(@transcriptions, &1)))
  end
end
