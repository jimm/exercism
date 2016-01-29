defmodule Allergies do
  use Bitwise

  @allergies %{
    1 => "eggs",
    2 => "peanuts",
    4 => "shellfish",
    8 => "strawberries",
    16 => "tomatoes",
    32 => "chocolate",
    64 => "pollen",
    128 => "cats"
  }
  @allergy_reverse_index %{
    "eggs" => 1,
    "peanuts" => 2,
    "shellfish" => 4,
    "strawberries" => 8,
    "tomatoes" => 16,
    "chocolate" => 32,
    "pollen" => 64,
    "cats" => 128
  }

  @doc """
  List the allergies for which the corresponding flag bit is true.
  """
  @spec list(non_neg_integer) :: [String.t]
  def list(flags), do: list(flags, 1, [])

  defp list(_, 256, allergies), do: allergies
  defp list(flags, bit, allergies) when band(flags, bit) == 0 do
    list(flags, bsl(bit, 1), allergies)
  end
  defp list(flags, bit, allergies) do
    list(flags, bsl(bit, 1), [@allergies[band(flags, bit)] | allergies])
  end

  @doc """
  Returns whether the corresponding flag bit in 'flags' is set for the item.
  """
  @spec allergic_to?(non_neg_integer, String.t) :: boolean
  def allergic_to?(0, item), do: false
  def allergic_to?(flags, item) do
    if band(flags, @allergy_reverse_index[item]) do
      true
    else
      false
    end
  end
end
