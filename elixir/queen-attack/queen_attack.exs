defmodule Queens do
  @type t :: %Queens{ black: {integer, integer}, white: {integer, integer} }
  defstruct black: nil, white: nil

  @doc """
  Creates a new set of Queens
  """
  @spec new(nil | list) :: Queens.t()
  def new(positions \\ %{white: {0, 3}, black: {7, 3}})

  def new(white: a, black: a), do: raise ArgumentError
  def new(positions),          do: %Queens{white: positions[:white], black: positions[:black]}

  @doc """
  Gives a string reprentation of the board with
  white and black queen locations shown
  """
  @spec to_string(Queens.t()) :: String.t()
  def to_string(queens) do
    (0..7) |> Enum.map_join("\n", &row_to_string(&1, queens))
  end

  defp row_to_string(row, queens) do
    (0..7) |> Enum.map_join(" ", &square_string(row, &1, queens))
  end

  defp square_string(row, col, %{black: {row, col}}), do: "B"
  defp square_string(row, col, %{white: {row, col}}), do: "W"
  defp square_string(_, _, _), do: "_"

  @doc """
  Checks if the queens can attack each other
  """
  @spec can_attack?(Queens.t()) :: boolean
  def can_attack?(%Queens{black: {row, _}, white: {row, _}}), do: true
  def can_attack?(%Queens{black: {_, col}, white: {_, col}}), do: true
  def can_attack?(%Queens{black: {b_row, b_col}, white: {w_row, w_col}}) do
    abs(w_row - b_row) == abs(w_col - b_col)
  end
end
