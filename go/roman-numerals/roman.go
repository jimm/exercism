package romannumerals

import "errors"

const testVersion = 2

var romanDigits = [][3]string{
	{"I", "V", "X"},
	{"X", "L", "C"},
	{"C", "D", "M"},
	{"M", "M", "M"},
}

func ToRomanNumeral(n int) (s string, err error) {
	if n <= 0 {
		err = errors.New("The Romans didn't know about non-whole numbers")
		return
	}
	if n >= 4000 {
		err = errors.New("The Romans couldn't count that high")
		return
	}

	place := 0
	for n > 0 {
		digit := n % 10
		places := romanDigits[place]
		s = digitToRoman(digit, places[0], places[1], places[2]) + s
		n = n / 10
		place++
	}
	return
}

func digitToRoman(digit int, low, med, high string) (s string) {
	switch digit {
	case 1:
		s = low
	case 2:
		s = low + low
	case 3:
		s = low + low + low
	case 4:
		s = low + med
	case 5:
		s = med
	case 6:
		s = med + low
	case 7:
		s = med + low + low
	case 8:
		s = med + low + low + low
	case 9:
		s = low + high
	case 0:
		s = ""
	}
	return
}
