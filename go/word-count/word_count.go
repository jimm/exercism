package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 2

type Frequency map[string]int

func WordCount(phrase string) (freqs Frequency) {
	freqs = make(Frequency)
	for _, word := range strings.Split(sanitize(phrase), " ") {
		if word != "" {
			freqs[strings.ToLower(word)] += 1
		}
	}
	return
}

func sanitize(s string) string {
	r_ := regexp.MustCompile("[^a-zA-Z0-9]")
	return string(r.ReplaceAll([]byte(s), []byte{' '}))
}
