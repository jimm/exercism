package gigasecond

import "time"

const TestVersion = 3
const gigasecond = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gigasecond)
}
