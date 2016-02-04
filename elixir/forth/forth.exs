defmodule Forth do
  @opaque evaluator :: any
  
  defstruct stack: [], words: %{}

  @doc """
  Create a new evaluator.
  """
  @spec new() :: evaluator
  def new() do
    %Forth{}
  end

  @doc """
  Evaluate an input string, updating the evaluator state.
  """
  @spec eval(evaluator, String.t) :: evaluator
  def eval(ev, s) do
    s
    |> String.downcase
    |> String.split(~r{[^-+\*/[:alnum:]]}u)
    |> Enum.reduce(ev, fn(word, ev) ->
      exec(ev, word)
    end)
  end

  @doc """
  Return the current stack as a string with the element on top of the stack
  being the rightmost element in the string.
  """
  @spec format_stack(evaluator) :: String.t
  def format_stack(%Forth{stack: stack}) do
    stack |> Enum.reverse |> Enum.join(" ")
  end

  # +
  defp exec(%Forth{stack: [a | [b | stack]]} = ev, "+") do
    val = (String.to_integer(b) + String.to_integer(a)) |> Integer.to_string
    %{ev | stack: [val | stack]}
  end
  defp exec(_, "+"), do: raise Forth.StackUnderflow

  # -
  defp exec(%Forth{stack: [a | [b | stack]]} = ev, "-") do
    val = (String.to_integer(b) - String.to_integer(a)) |> Integer.to_string
    %{ev | stack: [val | stack]}
  end
  defp exec(_, "-"), do: raise Forth.StackUnderflow

  # *
  defp exec(%Forth{stack: [a | [b | stack]]} = ev, "*") do
    val = (String.to_integer(b) * String.to_integer(a)) |> Integer.to_string
    %{ev | stack: [val | stack]}
  end
  defp exec(_, "*"), do: raise Forth.StackUnderflow

  # /
  defp exec(%Forth{stack: ["0" | _]}, "/"), do: raise Forth.DivisionByZero
  defp exec(%Forth{stack: [a | [b | stack]]} = ev, "/") do
    val = div(String.to_integer(b), String.to_integer(a)) |> Integer.to_string
    %{ev | stack: [val | stack]}
  end

  # dup
  defp exec(%Forth{stack: []}, "dup"), do: raise Forth.StackUnderflow
  defp exec(%Forth{stack: [a | _] = stack} = ev, "dup") do
    %{ev | stack: [a | stack]}
  end

  # drop
  defp exec(%Forth{stack: []}, "drop"), do: raise Forth.StackUnderflow
  defp exec(%Forth{stack: [a | stack]} = ev, "drop") do
    %{ev | stack: stack}
  end

  # swap
  defp exec(%Forth{stack: [a | [b | stack]]} = ev, "swap") do
    %{ev | stack: [b | [a | stack]]}
  end
  defp exec(_, "swap"), do: raise Forth.StackUnderflow

  # over
  defp exec(%Forth{stack: [a | [b | _]] = stack} = ev, "over") do
    %{ev | stack: [b | stack]}
  end
  defp exec(_, "over"), do: raise Forth.StackUnderflow

  # Everything else
  defp exec(%Forth{stack: stack} = ev, word) do
    %{ev | stack: [word | stack]}
  end

  # ================ exceptions ================

  defmodule StackUnderflow do
    defexception message: "stack underflow"
  end

  defmodule InvalidWord do
    defexception [:word]
    def exception(e), do: "invalid word: #{inspect e.word}"
    def message(e), do: exception(e)
  end

  defmodule UnknownWord do
    defexception [:word]
    def exception(e), do: "unknown word: #{inspect e.word}"
    def message(e), do: exception(e)
  end

  defmodule DivisionByZero do
    defexception message: "division by zero"
  end
end
