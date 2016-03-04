package meetup

import "time"

const testVersion = 3
const oneDay = time.Duration(24 * time.Hour)

const (
	First = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

type WeekSchedule int

func Day(sched WeekSchedule, day time.Weekday, month time.Month, year int) int {
	startDay := startDayFromSchedule(sched)
	if startDay >= 1 {
		t := time.Date(year, month, startDay, 0, 0, 0, 0, time.UTC)
		return findDay(t, day)
	}

	t := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	t = t.Add(-oneDay)
	return findDayFromEndOfMonth(t, day)
}

func startDayFromSchedule(sched WeekSchedule) int {
	switch sched {
	case First:
		return 1
	case Second:
		return 8
	case Third:
		return 15
	case Fourth:
		return 22
	case Teenth:
		return 13
	}
	return -1
}

func findDay(t time.Time, day time.Weekday) int {
	return findDayWithDirection(t, day, 1)
}

func findDayFromEndOfMonth(t time.Time, day time.Weekday) int {
	return findDayWithDirection(t, day, -1)
}

func findDayWithDirection(t time.Time, day time.Weekday, direction int) int {
	duration := oneDay
	if direction < 0 {
		duration = -oneDay
	}
	for t.Weekday() != day {
		t = t.Add(duration)
	}
	return t.Day()
}
