defmodule BinTree do
  @moduledoc """
  A node in a binary tree.

  `value` is the value of a node.
  `left` is the left subtree (nil if no subtree).
  `right` is the right subtree (nil if no subtree).

  A node in the zipper is a tuple of the form {BinTree, position} where
  position is either :left or :right (or nil for the root) and represents
  which child of the parent this node is.

  A zipper contains {curr_node, [parent_nodes]}.
  """
  @type t :: %BinTree{ value: any, left: BinTree.t | nil, right: BinTree.t | nil }
  defstruct value: nil, left: nil, right: nil
end

defmodule Zipper do
  @doc """
  Get a zipper focused on the root node.
  """
  @spec from_tree(BT.t) :: Z.t
  def from_tree(bt) do
    {{bt, nil}, []}
  end

  @doc """
  Get the complete tree from a zipper.
  """
  @spec to_tree(Z.t) :: BT.t
  def to_tree({{root, nil}, []}), do: root
  def to_tree(z), do: to_tree(up(z))

  @doc """
  Get the value of the focus node.
  """
  @spec value(Z.t) :: any
  def value(z) do
    {{node, _}, _} = z
    node.value
  end

  @doc """
  Get the left child of the focus node, if any.
  """
  @spec left(Z.t) :: Z.t | nil
  def left(z) do
    {{node, pos}, ancestors} = z
    if node.left do
      {{node.left, :left}, [{node, pos}|ancestors]}
    else
      nil
    end
  end
  
  @doc """
  Get the right child of the focus node, if any.
  """
  @spec right(Z.t) :: Z.t | nil
  def right(z) do
    {{node, pos}, ancestors} = z
    if node.right do
      {{node.right, :right}, [{node, pos}|ancestors]}
    else
      nil
    end
  end

  @doc """
  Get the parent of the focus node, if any.
  """
  @spec up(Z.t) :: Z.t
  def up(z) do
    {{node, position}, [{parent, parent_pos}|ancestors]} = z
    {{%{parent|position => node}, parent_pos}, ancestors}
  end

  @doc """
  Set the value of the focus node.
  """
  @spec set_value(Z.t, any) :: Z.t
  def set_value(z, v) do
    set(z, v, :value)
  end
  
  @doc """
  Replace the left child tree of the focus node. 
  """
  @spec set_left(Z.t, BT.t) :: Z.t
  def set_left(z, l) do
    set(z, l, :left)
  end
  
  @doc """
  Replace the right child tree of the focus node. 
  """
  @spec set_right(Z.t, BT.t) :: Z.t
  def set_right(z, r) do
    set(z, r, :right)
  end

  defp set(z, v, which) do
    {{node, position}, ancestors} = z
    {{%{node|which => v}, position}, ancestors}
  end
end
