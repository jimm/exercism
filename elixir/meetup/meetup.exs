defmodule Meetup do
  @moduledoc """
  Calculate meetup dates.
  """

  @weekday_map %{monday: 1, tuesday: 2, wednesday: 3, thursday: 4,
                 friday: 5, saturday: 6, sunday: 7}

  @type weekday ::
      :monday | :tuesday | :wednesday
    | :thursday | :friday | :saturday | :sunday

  @type schedule :: :first | :second | :third | :fourth | :last | :teenth

  @doc """
  Calculate a meetup date.

  The schedule is in which week (1..4, last or "teenth") the meetup date should
  fall.
  """
  @spec meetup(pos_integer, pos_integer, weekday, schedule) :: :calendar.date
  def meetup(year, month, weekday, schedule) do
    {year, month, find_date(year, month, Map.get(@weekday_map, weekday), schedule)}
  end

  defp find_date(year, month, weekday, :first) do
    find_date(year, month, weekday, 1)
  end
  defp find_date(year, month, weekday, :second) do
    find_date(year, month, weekday, 8)
  end
  defp find_date(year, month, weekday, :third) do
    find_date(year, month, weekday, 15)
  end
  defp find_date(year, month, weekday, :fourth) do
    find_date(year, month, weekday, 22)
  end
  defp find_date(year, month, weekday, :teenth) do
    find_date(year, month, weekday, 13)
  end
  defp find_date(year, month, weekday, :last) do
    last_day = Enum.drop_while((31..28), fn(i) ->
      not :calendar.valid_date(year, month, i)
    end)
    |> Enum.take(1)
    |> hd
    Enum.drop_while((last_day..(last_day-7)), fn(i) ->
      :calendar.day_of_the_week(year, month, i) != weekday
    end)
    |> Enum.take(1)
    |> hd
  end

  defp find_date(year, month, weekday, start) do
    Enum.drop_while((start..start+7), fn(i) ->
      :calendar.day_of_the_week(year, month, i) != weekday
    end)
    |> Enum.take(1)
    |> hd
  end
end
