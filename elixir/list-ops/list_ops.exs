defmodule ListOps do
  # Please don't use any external modules (especially List) in your
  # implementation. The point of this exercise is to create these basic functions
  # yourself.
  #
  # Note that `++` is a function from an external module (Kernel, which is
  # automatically imported) and so shouldn't be used either.

  @spec count(list) :: non_neg_integer
  def count(l), do: count(l, 0)
  defp count([], len), do: len
  defp count([_|t], len), do: count(t, len+1)

  @spec reverse(list) :: list
  def reverse(l), do: reverse(l, [])
  defp reverse([], rev), do: rev
  defp reverse([h|t], rev), do: reverse(t, [h|rev])

  @spec map(list, (any -> any)) :: list
  def map(l, f), do: map(l, f, [])
  defp map([], _, acc), do: reverse(acc)
  defp map([h|t], f, acc), do: map(t, f, [f.(h)|acc])

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(l, f), do: filter(l, f, [])
  defp filter([], _, acc), do: reverse(acc)
  defp filter([h|t], f, acc) do
    if f.(h) do
      filter(t, f, [h|acc])
    else
      filter(t, f, acc)
    end
  end

  @type acc :: any
  @spec reduce(list, acc, ((any, acc) -> acc)) :: acc
  def reduce([], acc, _), do: acc
  def reduce([h|t], acc, f), do: reduce(t, f.(h, acc), f)

  @spec append(list, list) :: list
  def append([], []), do: []
  def append([], b), do: b
  def append(a, []), do: a
  def append(a, b), do: do_append(reverse(a), b)
  defp do_append([], b), do: b
  defp do_append([h|t], b), do: do_append(t, [h|b])

  @spec concat([[any]]) :: [any]
  def concat(ll), do: concat(ll, [])
  defp concat([], acc), do: reverse(acc)
  defp concat([h|t], acc), do: concat(t, append(reverse(h), acc))
end
