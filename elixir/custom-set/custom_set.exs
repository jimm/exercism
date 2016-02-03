# I could use a `Map` whose keys are the values in the `Set`, but that feels
# like cheating because it'd make everything so easy.
defmodule CustomSet do
  # This lets the compiler check that all Set callback functions have been
  # implemented.
  @behaviour Set

  defstruct elements: []

  def new() do
    %CustomSet{}
  end

  def new(coll) do
    %CustomSet{elements: coll |> Enum.uniq}
  end

  def empty(), do: new

  def delete(cs, value) do
    deleted = cs.elements |> Enum.reduce([], fn(v, acc) ->
      if v === value, do: acc, else: [v | acc]
    end)
    %{cs | elements: deleted}
  end

  def difference(cs1, cs2) do
    # Enum.member? uses ===. The tests require == (for example, 2 must match
    # 2.0) so I can't use member?.
    diff = cs1.elements |> Enum.reduce([], fn(v, acc) ->
      found = cs2.elements |> Enum.reduce(false, fn(v2, acc2) ->
        if v == v2, do: true, else: acc2
      end)
      if found, do: acc, else: [v|acc]
     end)
    IO.puts "diff of #{inspect cs1.elements} and #{inspect cs2.elements} = #{inspect Enum.reverse(diff)}"
    %CustomSet{elements: Enum.reverse(diff)}
  end

  def disjoint?(cs1, cs2) do
    # We could make this faster by short-circuiting, but with the sizes
    # we're working with it doesn't matter.
    same = for v1 <- cs1.elements,
               v2 <- cs2.elements,
               v1 === v2 do
                 v1
               end
    length(same) == 0
  end

  def equal?(%CustomSet{elements: coll1}, %CustomSet{elements: coll2}) when length(coll1) != length(coll2), do: false
  def equal?(%CustomSet{elements: coll1}, %CustomSet{elements: coll2}) do
    Enum.zip(Enum.sort(coll1), Enum.sort(coll2))
    |> Enum.reduce(true, fn({v1, v2}, acc) ->
      if v1 != v2, do: false, else: acc
    end)
  end

  def intersection(cs1, cs2) do
    cs1
  end

  def member?(cs, value) do
    true
  end

  def put(cs, value) do
    cs
  end

  def size(cs) do
    length(cs.elements)
  end

  def subset?(cs1, cs2) do
    true
  end

  def to_list(%{elements: es}) do
    es
  end

  def union(cs1, cs2) do
    cs1
  end
end
