package igpay

// TODO peel off front instead of hard-coding list of dipthongs
// TDOO include 'u' if preceeded by 'q'
// TODO handle words

func PigLatin(s string) string {
	dipthong := maybeDipthong(s[0:2])
	if dipthong != "" {
		return s[2:] + dipthong + "ay"
	} else if isConsonant(s[0:1]) {
		return s[1:] + s[0:1] + "ay"
	} else {
		return s + "ay"
	}
}

func maybeDipthong(s string) string {
	dipthong := s[0:2]
	if dipthong == "st" || dipthong == "qu" || dipthong == "th" || dipthong == "ch" {
		return dipthong
	}
	return ""	
}

func isConsonant(s string) bool {
	return s != "a" && s != "e" && s != "i" && s != "o" && s != "u" && s != "y"
}
