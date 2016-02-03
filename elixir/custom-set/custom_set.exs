# I could use a `Map` whose keys are the values in the `Set`, but that feels
# like cheating because it'd make everything so easy.
defmodule CustomSet do
  # This lets the compiler check that all Set callback functions have been
  # implemented.
  @behaviour Set

  defstruct list: []

  def new() do
    %CustomSet{}
  end

  def new(coll) do
    %CustomSet{list: coll |> Enum.uniq |> Enum.sort}
  end

  def empty(_cs), do: new

  def delete(cs, value) do
    new(cs.list |> List.delete(value))
  end

  def difference(cs1, cs2) do
    new(cs1.list -- cs2.list)
  end

  def disjoint?(cs1, cs2) do
    # We could make this faster by short-circuiting, but with the sizes
    # we're working with it doesn't matter.
    same = for v1 <- cs1.list,
               v2 <- cs2.list,
               v1 === v2 do
                 v1
               end
    length(same) == 0
  end

  def equal?(%CustomSet{list: s1}, %CustomSet{list: s2}) when length(s1) != length(s2), do: false
  def equal?(%CustomSet{list: s1}, %CustomSet{list: s2}) do
    Enum.zip(Enum.sort(s1), Enum.sort(s2))
    |> Enum.reduce(true, fn({v1, v2}, acc) ->
      if v1 != v2, do: false, else: acc
    end)
  end

  def intersection(cs1, cs2) do
    same = for v1 <- cs1.list,
               v2 <- cs2.list,
               v1 === v2 do
                 v1
               end
    new(same)
  end

  def member?(cs, value) do
    length(cs.list) != length(cs.list -- [value])
  end

  def put(cs, value) do
    new([value | cs.list] |> Enum.sort)
  end

  def size(cs), do: length(cs.list)

  def subset?(cs1, cs2) do
    length(cs1.list -- cs2.list) == 0
  end

  def to_list(cs), do: cs.list

  def union(cs1, cs2), do: new(cs1.list ++ cs2.list)
end
