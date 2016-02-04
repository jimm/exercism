defmodule Forth do
  @opaque evaluator :: any
  
  defstruct stack: [], words: %{}, defining: nil

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
    |> words
    |> Enum.reduce(ev, fn(word, ev) -> eval_word(ev, word) end)
  end

  @doc """
  Return the current stack as a string with the element on top of the stack
  being the rightmost element in the string.
  """
  @spec format_stack(evaluator) :: String.t
  def format_stack(%Forth{stack: stack}) do
    stack |> Enum.reverse |> Enum.join(" ")
  end

  defp eval_word(ev, word) do
    definition = Map.get(ev.words, word, nil)
    if ev.defining == nil && definition do
      eval(ev, Enum.join(definition, " "))
    else
      exec(ev, word)
    end
  end

  # ================ defining new words ================

  # :, set defining to true (will become name of definition when seen)
  defp exec(%Forth{defining: nil} = ev, ":") do
    %{ev | defining: true}
  end

  # looking for name of word to define
  defp exec(%Forth{defining: true} = ev, name) do
    if String.match?(name, ~r{\A[[:digit:]]}) do
      raise Forth.InvalidWord
    else
      %{ev | defining: name, words: Map.put(ev.words, name, [])}
    end
  end

  # ;
  defp exec(%Forth{defining: name} = ev, ";") when name != nil and name != true do
    %{ev | defining: nil}
  end

  # while defining a new word
  defp exec(%Forth{defining: name} = ev, word) when name != nil do
    %{ev | words: Map.put(ev.words, name,
                          [word | Map.get(ev.words, name)])}
  end

  # ================ mathematics ================

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

  # ================ built-in words ================

  # dup
  defp exec(%Forth{stack: []}, "dup"), do: raise Forth.StackUnderflow
  defp exec(%Forth{stack: [a | _] = stack} = ev, "dup") do
    %{ev | stack: [a | stack]}
  end

  # drop
  defp exec(%Forth{stack: []}, "drop"), do: raise Forth.StackUnderflow
  defp exec(%Forth{stack: [_ | stack]} = ev, "drop") do
    %{ev | stack: stack}
  end

  # swap
  defp exec(%Forth{stack: [a | [b | stack]]} = ev, "swap") do
    %{ev | stack: [b | [a | stack]]}
  end
  defp exec(_, "swap"), do: raise Forth.StackUnderflow

  # over
  defp exec(%Forth{stack: [_ | [b | _]] = stack} = ev, "over") do
    %{ev | stack: [b | stack]}
  end
  defp exec(_, "over"), do: raise Forth.StackUnderflow

  # ================ everything else ================

  defp exec(%Forth{stack: stack} = ev, word) do
    if String.match?(word, ~r{\A[[:digit:]]+\z}) do
      %{ev | stack: [word | stack]}
    else
      raise Forth.UnknownWord
    end
  end

  # ================ helpers ================

  defp words(s) do
    s
    |> String.downcase
    |> String.split(~r{[\s[:cntrl:]]}u)
  end

  # ================ exceptions ================

  defmodule StackUnderflow do
    defexception message: "stack underflow"
  end

  defmodule InvalidWord do
    defexception message: "invalid word"

    # I can't get the below to work.

    # defexception [:word]
    # def exception(e), do: "invalid word: #{inspect e.word}"
    # def message(e), do: exception(e)
  end

  defmodule UnknownWord do
    defexception message: "unknown word"

    # I can't get the below to work.

    # defexception [:word]
    # def exception(e), do: "unknown word: #{inspect e.word}"
    # def message(e), do: exception(e)
  end

  defmodule DivisionByZero do
    defexception message: "division by zero"
  end
end
