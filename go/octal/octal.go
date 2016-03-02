package octal

import "errors"

func ParseOctal(s string) (int64, error) {
	var n int64
	for _, r := range s {
		if r >= '0' && r <= '7' {
			n = n*8 + int64(r-'0')
		} else {
			return 0, errors.New("wha?")
		}
	}
	return n, nil
}
