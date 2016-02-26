defmodule RunLengthEncoder do 

  @doc """
  Generates a string where consecutive elements are represented as a data value and count.
  "HORSE" => "1H1O1R1S1E"
  For this example, assume all input are strings, that are all uppercase letters.
  It should also be able to reconstruct the data into its original form. 
  "1H1O1R1S1E" => "HORSE" 
  """
  # @spec encode(string) :: String.t
  def encode(string) do
    encode(string, 0, 0, "")
  end

  def encode("", 0, _, _), do: ""
  def encode("", n, c, answer), do: add_encoded(answer, n, c)
  def encode(<<c::utf8, rest::binary>>, n, c, answer) do
    encode(rest, n+1, c, answer)
  end
  def encode(<<c::utf8, rest::binary>>, 0, _, answer) do
    encode(rest, 1, c, answer)
  end
  def encode(<<c::utf8, rest::binary>>, n, prev_c, answer) do
    encode(rest, 1, c, add_encoded(answer, n, prev_c))
  end

  defp add_encoded(answer, n, c), do: "#{answer}#{n}#{<<c>>}"

  # @spec decode(string) :: String.t
  def decode(string) do
    decode(string, "")
  end

  def decode("", answer), do: answer
  def decode(string, answer) do
    [s, num_str, letter] = Regex.run(~r{(\d+)([A-Z])}, string)

    num = String.to_integer(num_str)
    repeated_chars = List.duplicate(letter, num)
    |> Enum.join

    decode(String.slice(string, String.length(s),
                        String.length(string) - String.length(s)),
           "#{answer}#{repeated_chars}")
  end
end
