defmodule Minesweeper do

  @doc """
  Annotate empty spots next to mines with the number of mines next to them.
  """
  @spec annotate([String.t]) :: [String.t]

  def annotate([]), do: []
  def annotate(board) do
    num_rows = length(board)
    num_cols = String.length(hd(board))
    board
    |> to_array
    |> insert_numbers(num_rows, num_cols)
    |> to_strings(num_cols)
  end

  defp to_array(board) do
    board |> Enum.flat_map(&(String.graphemes(&1)))
  end

  defp insert_numbers(array, num_rows, num_cols) do
    (0..length(array) - 1)
    |> Enum.map(&(number_or_mine(array, &1, num_rows, num_cols)))
  end

  defp number_or_mine(array, n, num_rows, num_cols) do
    if Enum.at(array, n) == "*" do
      "*"
    else
      number(array, n, num_rows, num_cols)
    end
  end

  defp number(array, n, num_rows, num_cols) do
    col = rem(n, num_cols)
    row = div((n-col), num_cols)
    count = Enum.sum(for r <- (row-1..row+1),
                         c <- (col-1..col+1),
                         r >= 0 && r < num_rows &&
                         c >= 0 && c < num_cols &&
                         !(r == row && c == col) &&
                         Enum.at(array, r * num_cols + c) == "*" do
                           1
                         end)
    if count == 0, do: " ", else: count
  end

  defp to_strings(array, num_cols) do
    array |> Enum.chunk(num_cols) |> Enum.map(&(Enum.join(&1)))
  end
end