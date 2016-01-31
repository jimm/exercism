defmodule Graph do
  defstruct attrs: [], nodes: [], edges: []
end

defmodule Dot do
  defmacro graph([do: nil]) do
    quote do
      %Graph{}
    end
  end
  defmacro graph([do: block]) do
    IO.inspect block
    do_graph(block)
  end

  def do_graph(lines) when is_list(lines) do
    graph = Enum.reduce(lines, %Graph{}, fn(line, g) ->
      line_graph = do_graph(line)
      g = %{g | attrs: g.attrs ++ line_graph.attrs}
      g = %{g | nodes: g.nodes ++ line_graph.nodes}
      %{g | edges: g.edges ++ line_graph.edges}
    end)
    quote do
      unquote(graph)
    end
  end
  def do_graph({:graph, _, attrs}) do
    quote do
      %Graph{attrs: unquote(hd(attrs))}
    end
  end
  def do_graph({:--, _, [arg1 | [arg2 | attrs]]}) do
    g1 = do_graph(arg1)
    g2 = do_graph(arg2)
    quote do
      %Graph{edges: [{unquote(g1.nodes[0]), unquote(g2.nodes[0]), unquote(attrs)}]}
    end
  end
  def do_graph({node, _, args}) when is_atom(node) do
    attrs = if args do
      hd(args)
    else
      []
    end
    quote do
      %Graph{nodes: [{unquote(node), unquote(attrs)}]}
    end
  end
  def do_graph({node, _, _}) when is_list(node) do
    quote do
      %Graph{attrs: unquote(node)}
    end
  end
  def do_graph(_) do
    raise ArgumentError
  end
end
