package trinary

import "errors"

func ParseTrinary(s string) (int64, error) {
	var val int64
	for _, r := range s {
		switch r {
		case '0':
			val = val * 3
		case '1':
			val = val*3 + 1
		case '2':
			val = val*3 + 2
		default:
			return 0, errors.New("huh?")
		}
		if val < 0 {
			return 0, errors.New("overflow")
		}
	}
	return val, nil
}
