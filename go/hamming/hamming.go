package hamming

import "errors"

const TestVersion = 2

func Distance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("strings must be of equal length")
	}
	dist := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			dist += 1
		}
	}
	return dist, nil
}
