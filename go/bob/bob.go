package bob

import "strings"

const testVersion = 2

func Hey(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Fine. Be that way!"
	}
	if strings.ToUpper(s) == s && strings.ToLower(s) != s {
		return "Whoa, chill out!"
	}
	if s[len(s)-1] == '?' {
		return "Sure."
	}
	return "Whatever."
}
