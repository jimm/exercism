package leap

const TestVersion = 1

func IsLeapYear(year int) (is bool) {
	switch {
	case year%400 == 0:
		is = true
	case year%100 == 0:
		is = false
	case year%4 == 0:
		is = true
	}
	return
}
