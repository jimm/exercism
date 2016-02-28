package atbash

import (
	"strings"
	"unicode"
)

const splitLen = 5

var atbash = map[rune]rune{
	'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u',
	'g': 't', 'h': 's', 'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o',
	'm': 'n', 'n': 'm', 'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i',
	's': 'h', 't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c',
	'y': 'b', 'z': 'a',
}

func Atbash(plaintext string) string {
	s := []rune{}
	for _, c := range strings.ToLower(plaintext) {
		atbash := atbash[c]
		if atbash != 0 {
			s = append(s, atbash)
		} else if unicode.IsNumber(c) {
			s = append(s, c)
		}
	}
	return strings.Join(splitIntoGroups(string(s)), " ")
}

func splitIntoGroups(s string) (strs []string) {
	for len(s) > 0 {
		if splitLen > len(s) {
			strs = append(strs, s)
			s = ""
		} else {
			strs = append(strs, s[0:splitLen])
			s = s[splitLen:]
		}
	}
	return
}
