package say

import "strings"

var powers = []string{
	"",
	"thousand",
	"million",
	"trillion",
	"billion",
	"quadrillion",
	"quintillion",
	"sextillion",
	"septillion",
	"octillion",
	"nonillion",
	"decillion"
}

func Say(n uint64) string {
	if n == 0 {
		return "zero"
	}
	words := make([]string)
	power := 0
	while n > 0 {
		words = sayWithPower(n % 100, powers[power], words)
		power++
	}
	return strings.Join(words, " ")
}

func sayWithPower(n uint, powerWord string, words []string) []string {
	words = sayNumber(n, words)
}
