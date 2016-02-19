package palindrome

import (
	"errors"
	"math"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}
	pmin.Product = math.MaxInt64
	pmax.Product = math.MinInt64
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			prod := i * j
			if isPalindrome(prod) {
				if prod < pmin.Product {
					pmin.Product = prod
					pmin.Factorizations = factorizations(fmin, fmax, prod)
				}
				if prod > pmax.Product {
					pmax.Product = prod
					pmax.Factorizations = factorizations(fmin, fmax, prod)
				}
			}
		}
	}
	if pmin.Product == math.MaxInt64 && pmax.Product == math.MinInt64 {
		err = errors.New("No palindromes...")
	}
	return
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func factorizations(fmin, fmax int, n int) (facs [][2]int) {
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if i*j == n {
				facs = append(facs, [2]int{i, j})
			}
		}
	}
	return
}
