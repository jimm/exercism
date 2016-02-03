defmodule Graph do
  defstruct attrs: [], nodes: [], edges: []
end

defmodule Dot do

  defmacro graph([do: {:__block__, _, lines}]) when is_list(lines) do
    Enum.reduce(lines, %Graph{}, fn(line, g) ->
      build_graph(g, line)
    end)
    |> Macro.escape
  end
  defmacro graph([do: block]) do
    build_graph(%Graph{}, block)
    |> Macro.escape
  end

  def build_graph(g, nil), do: g
  def build_graph(g, {:graph, _, [[]]}), do: g
  def build_graph(g, {:graph, _, attrs}) do
    # The tests want the attrs sorted by the first element
    attrs = [hd(hd(attrs)) | g.attrs]
    |> Enum.sort
    %{g | attrs: attrs}
  end
  def build_graph(g, {:--, _, [arg1 | [arg2 | _]]}) do
    g1 = build_graph(%Graph{}, arg1)
    g2 = build_graph(%Graph{}, arg2)
    [{n1, _}] = g1.nodes
    [{n2, attrs}] = g2.nodes
    edges = [{n1, n2, attrs} | g.edges]
    |> Enum.sort
    %{g | edges: edges}
  end
  def build_graph(g, {node, _, nil}) when is_atom(node) do
    nodes = [{node, []} | g.nodes]
    |> Enum.sort
    %{g | nodes: nodes}
  end
  def build_graph(g, {node, line_info, [[]]}) when is_atom(node) do
    build_graph(g, {node, line_info, nil})
  end
  def build_graph(g, {node, _, [[{key, val}]]}) when is_atom(node) do
    nodes = [{node, [{key, val}]} | g.nodes]
    |> Enum.sort
    %{g | nodes: nodes}
  end
  def build_graph(g, {attrs, _, _}) when is_list(attrs) do
    new_attrs = [hd(attrs) | g.attrs]
    |> Enum.sort
    %{g | attrs: new_attrs}
  end
  def build_graph(_, _) do
    raise ArgumentError
  end
end
