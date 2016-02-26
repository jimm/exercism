defmodule Acronym do
  @doc """
  Generate an acronym from a string. 
  "This is a string" => "TIAS"
  """
  # @spec abbreviate(string) :: String.t()
  def abbreviate(string) do
    string
    |> String.split(~r{[-\s]})
    |> Enum.flat_map(&split_on_case/1)
    |> Enum.map(&(String.slice(&1, 0, 1)))
    |> Enum.join("")
    |> String.upcase
  end

  def split_on_case(word) do
    words = Regex.scan(~r{[A-Z][a-z]*}, word)
    |> Enum.map(&hd/1)
    if Enum.empty?(words), do: [word], else: words
  end
end
