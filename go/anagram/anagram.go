package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (anagrams []string) {
	lSubject := strings.ToLower(subject)
	sortedSubject := sortedChars(lSubject)
	for _, s := range candidates {
		ls := strings.ToLower(s)
		if equalChars(sortedSubject, sortedChars(ls)) && lSubject != ls {
			anagrams = append(anagrams, ls)
		}
	}
	return
}

func sortedChars(s string) []string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return chars
}

func equalChars(ss1 []string, ss2 []string) bool {
	if len(ss1) != len(ss2) {
		return false
	}
	for i, c1 := range ss1 {
		if c1 != ss2[i] {
			return false
		}
	}
	return true
}
