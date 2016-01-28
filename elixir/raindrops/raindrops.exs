defmodule Raindrops do

  @raindrop_factors [{3, "Pling"}, {5, "Plang"}, {7, "Plong"}]

  @doc """
  Returns a string based on raindrop factors.

  - If the number contains 3 as a prime factor, output 'Pling'.
  - If the number contains 5 as a prime factor, output 'Plang'.
  - If the number contains 7 as a prime factor, output 'Plong'.
  - If the number does not contain 3, 5, or 7 as a prime factor,
    just pass the number's digits straight through.
  """
  @spec convert(pos_integer) :: String.t
  def convert(number) do
    strs = @raindrop_factors
    |> Enum.map(fn({divisor, str}) -> factor_word(number, divisor, str) end)
    case strs do
      ["", "", ""] -> Integer.to_string(number)
      [a, b, c] -> a <> b <> c
    end
  end

  defp factor_word(number, divisor, str) when rem(number, divisor) == 0, do: str
  defp factor_word(_, _, _), do: ""
end
