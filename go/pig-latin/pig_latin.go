package igpay

import "strings"

func PigLatin(s string) string {
	words := strings.Split(s, " ")
	igpayAtinlay := make([]string, 0)
	for _, word := range words {
		igpayAtinlay = append(igpayAtinlay, pigLatin(word))
	}
	return strings.Join(igpayAtinlay, " ")
}

func pigLatin(s string) string {
	front := leadingConsonants(s)
	if front != "" {
		return s[len(front):] + front + "ay"
	}
	return s + "ay"
}

func leadingConsonants(s string) string {
	front := make([]rune, 0)
	for i, r := range s {
		if consonantSound(i, r, s) {
			front = append(front, r)
		} else {
			break
		}
	}
	if len(front) == 0 {
		return s
	}
	if front[len(front)-1] == 'q' && s[len(front):len(front)+1] == "u" {
		front = append(front, 'u')
	}
	return string(front)
}

func consonantSound(i int, r rune, s string) bool {
	if isVowel(r) {
		return false
	}
	if isVowelIfFollwedByVowel(r) && !isVowel(rune(s[i+1])) {
		return false
	}
	return true
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func isVowelIfFollwedByVowel(r rune) bool {
	return r == 'x' || r == 'y'
}
