defmodule Sublist do
  @doc """
  Returns whether the first list is a sublist or a superlist of the second list
  and if not whether it is equal or unequal to the second list.
  """
  def compare(a, a), do: :equal
  def compare([], [_|_]), do: :sublist
  def compare([_|_], []), do: :superlist
  def compare(a, b) when length(a) > length(b) do
    compare(b, a, length(b), 0, length(a) - length(b) + 1, :superlist)
  end
  def compare(a, b) when length(a) < length(b) do
    compare(a, b, length(a), 0, length(b) - length(a) + 1, :sublist)
  end
  def compare(_, _), do: :unequal

  # At this point, a is shorter than b. If a is found in b, return
  # sub_or_super. Else return :unequal.
  def compare(_, _, _, beyond_end_index, beyond_end_index, _), do: :unequal
  def compare(a, b, len_a, index, beyond_end_index, sub_or_super) do
    if a === Enum.slice(b, index, len_a) do
      sub_or_super
    else
      compare(a, b, len_a, index+1, beyond_end_index, sub_or_super)
    end
  end
end
