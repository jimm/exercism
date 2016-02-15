package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	hour   int
	minute int
}

// Time returns a Clock given an hour and minute.
func Time(hour, minute int) Clock {
	return Clock{hour, minute}.normalize()
}

// String return "hh:mm", given a Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds mintues to a Clock.
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.normalize()
}

// normalize makes sure minutes are 0..60 and hour is 0..23.
func (c Clock) normalize() Clock {
	for ; c.minute >= 60; c.minute -= 60 {
		c.hour++
	}
	for ; c.minute < 0; c.minute += 60 {
		c.hour--
	}
	for ; c.hour >= 24; c.hour -= 24 {
	}
	for ; c.hour < 0; c.hour += 24 {
	}
	return c
}
