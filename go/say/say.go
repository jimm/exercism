package say

import "fmt"					// DEBUG
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
		fmt.Println("Say n", n)
		words = sayWithPower(n % 100, powers[power], words)
		n /= 100
		power++
	}
	fmt.Println("finally, words before join", words)
	return strings.Join(words, " ")
}

func sayWithPower(n uint64, powerWord string, words []string) []string {
	fmt.Println("sayWithPower", n, powerWord, words)
	words = sayNumber(n, words)
	if powerWord != "" {
		words = append(words, powerWord)
	}
	return words
}

func sayNumber(n uint64, words []string) []string {
	fmt.Println("sayNumber", n, words)
	if n < 10 || n >= 20 {
		word := ""
		tensPlace := n / 10
		if tens[tensPlace] != "" {
			word += tens[tensPlace]
		}
		digit := n%10
		if digit > 0 && tens[tensPlace] != "" {
			word += "-" + digits[digit]
		}
		words = append(words, word)
	} else {
		words = append(words, teens[n-10])
	}
	return words
}
