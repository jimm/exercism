package say

import "strings"

var powers = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion",
	"sextillion",
	"septillion",
	"octillion",
	"nonillion",
	"decillion",
}
var digits = []string{
	"",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}
var teens = []string{
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}
var tens = []string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func Say(n uint64) string {
	if n == 0 {
		return "zero"
	}
	words := make([]string, 0)
	power := 0
	for n > 0 {
		nextWord := sayWithPower(n%1000, powers[power])
		if nextWord != "" {
			words = append([]string{nextWord}[0:], words...)
		}
		n /= 1000
		power++
	}
	return strings.TrimSpace(strings.Join(words, " "))
}

func sayWithPower(n uint64, powerWord string) string {
	if n == 0 {
		return ""
	}
	word := sayNumber(n)
	if powerWord != "" {
		word += " " + powerWord
	}
	return word
}

func sayNumber(n uint64) string {
	word := sayHundreds(n)
	n = n % 100
	if n > 0 {
		if word != "" {
			word += " "
		}
		word += sayTo99(n)
	}
	return word
}

func sayHundreds(n uint64) string {
	if n >= 100 {
		return sayTo99(n/100) + " hundred"
	}
	return ""
}

func sayTo99(n uint64) string {
	if n < 10 {
		return digits[n]
	} else if n < 20 {
		return teens[n-10]
	}

	tensPlace := (n % 100) / 10
	onesPlace := n % 10

	word := ""
	if tens[tensPlace] != "" {
		word += tens[tensPlace]
	}
	if onesPlace > 0 {
		if tens[tensPlace] != "" {
			word += "-" + digits[onesPlace]
		} else {
			if word != "" {
				word += " "
			}
			word += digits[onesPlace]
		}
	}
	return word
}
