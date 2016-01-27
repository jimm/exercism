defmodule Gigasecond do
	@doc """
	Calculate a date one billion seconds after an input date.
	"""
	@spec from({{pos_integer, pos_integer, pos_integer}, {pos_integer, pos_integer, pos_integer}}) :: :calendar.datetime

	def from({{year, month, day}, {hours, minutes, seconds}} = dt) do
    secs = :calendar.datetime_to_gregorian_seconds(dt)
    :calendar.gregorian_seconds_to_datetime(secs + 1000000000)
	end
end
