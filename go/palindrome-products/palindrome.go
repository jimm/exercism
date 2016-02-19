package palindrome

import (
	// "errors"
	"math"
	"strconv"
	// "strings"
)

type Product int

func Products(min, max int) (minProd, maxProd int, err error) {
	minProd = math.MaxInt64
	maxProd = math.MinInt64
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			prod := i * j
			if isPalindrome(prod) {
				if prod < minProd {
					minProd = prod
				}
				if prod > maxProd {
					maxProd = prod
				}
			}
		}
	}
	return
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s) / 2 ; i++ {
		if s[i] != s[len(s) - 1 - i] {
			return false;
		}
	}
	return true
}
