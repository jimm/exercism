package etl

import "strings"

func Transform(old map[int][]string) map[string]int {
	m := make(map[string]int)
	for score, letters := range old {
		for _, letter := range letters {
			m[strings.ToLower(letter)] = score
		}
	}
	return m
}
