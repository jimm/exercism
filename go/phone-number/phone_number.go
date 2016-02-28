package phonenumber

import (
	"errors"
	"unicode"
)

func Number(s string) (string, error) {
	s = onlyDigits(s)
	s, err := checkLength(s)
	if err != nil {
		return "", err
	}
	return s, nil
}

func AreaCode(s string) (string, error) {
	num, err := Number(s)
	if err != nil {
		return "", err
	}
	return num[0:3], nil
}

func Format(s string) (string, error) {
	num, err := checkLength(s)
	if err != nil {
		return "", err
	}
	return "(" + num[0:3] + ") " + num[3:6] + "-" + num[6:], nil
}

func onlyDigits(s string) string {
	runes := []rune{}
	for _, r := range s {
		if unicode.IsDigit(r) {
			runes = append(runes, r)
		}
	}
	return string(runes)
}

func checkLength(s string) (string, error) {
	if len(s) < 10 {
		return "", errors.New("not enough digits")
	}
	if len(s) == 11 {
		if s[0] == '1' {
			s = s[1:]
		} else {
			return "", errors.New("11-digit number must start with 1")
		}
	}
	if len(s) > 11 {
		return "", errors.New("too many cooks, I mean digits")
	}
	return s, nil
}
