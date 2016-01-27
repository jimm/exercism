defmodule Roman do
  @doc """
  Convert the number to a roman number.
  """
  @spec numerals(pos_integer) :: String.t
  def numerals(number), do: numerals(number, 1, [""])

  defp numerals(0, _, acc), do: acc |> Enum.join("")
  defp numerals(number, 1, acc) do
    {rest, digit} = lsdigit(number)
    roman = digit_to_roman(digit, "I", "V", "X")
    numerals(rest, 2, [roman|acc])
  end
  defp numerals(number, 2, acc) do
    {rest, digit} = lsdigit(number)
    roman = digit_to_roman(digit, "X", "L", "C")
    numerals(rest, 3, [roman|acc])
  end
  defp numerals(number, 3, acc) do
    {rest, digit} = lsdigit(number)
    roman = digit_to_roman(digit, "C", "D", "M")
    numerals(rest, 4, [roman|acc])
  end
  defp numerals(number, 4, acc) do
    {rest, digit} = lsdigit(number)
    roman = case digit do
              1 -> "M"
              2 -> "MM"
              3 -> "MMM"
              _ -> ""           # huh? numbers bigger than 3999?
            end
    numerals(rest, 5, [roman|acc])
  end

  defp lsdigit(number) do
    rest = div(number, 10)
    digit = number - (rest * 10)
    {rest, digit}
  end

  defp digit_to_roman(digit, low, med, high) do
    case digit do
      1 -> "#{low}"
      2 -> "#{low}#{low}"
      3 -> "#{low}#{low}#{low}"
      4 -> "#{low}#{med}"
      5 -> "#{med}"
      6 -> "#{med}#{low}"
      7 -> "#{med}#{low}#{low}"
      8 -> "#{med}#{low}#{low}#{low}"
      9 -> "#{low}#{high}"
      0 -> ""
    end
  end
end
