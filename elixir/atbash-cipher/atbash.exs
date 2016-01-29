defmodule Atbash do

  @atbash %{
    "a" => "z", "b" => "y", "c" => "x", "d" => "w", "e" => "v", "f" => "u",
    "g" => "t", "h" => "s", "i" => "r", "j" => "q", "k" => "p", "l" => "o",
    "m" => "n", "n" => "m", "o" => "l", "p" => "k", "q" => "j", "r" => "i",
    "s" => "h", "t" => "g", "u" => "f", "v" => "e", "w" => "d", "x" => "c",
    "y" => "b", "z" => "a"
  }

  @doc """
  Encode a given plaintext to the corresponding ciphertext

  ## Examples

  iex> Atbash.encode("completely insecure")
  "xlnko vgvob rmhvx fiv"
  """
  @spec encode(String.t) :: String.t
  def encode(plaintext) do
    plaintext
    |> String.downcase
    |> String.replace(~r{[^a-z0-9]}, "")
    |> String.graphemes
    |> Enum.map(&(@atbash[&1] || &1))
    |> split_into_groups_of(5, [])
    |> Enum.join(" ")
  end

  defp split_into_groups_of([], _, acc), do: Enum.reverse(acc)
  defp split_into_groups_of(strs, n, acc) do
    split_into_groups_of(Enum.drop(strs, n), n,
                         [Enum.take(strs, n) |> Enum.join("") | acc])
  end
end
