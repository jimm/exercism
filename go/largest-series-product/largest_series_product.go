package lsproduct

import "errors"

const TestVersion = 2

func LargestSeriesProduct(s string, span int) (int, error) {
	if span > len(s) {
		return 0, errors.New("span must be <= length of string")
	}
	if span < 0 {
		return 0, errors.New("span must be >= 0")
	}

	if span == 0 {
		return 1, nil
	}

	digits, err := stringToDigits(s)
	if err != nil {
		return 0, err
	}

	maxProd := 0
	for i := 0; i < len(s)-span+1; i++ {
		prod := mult(digits[i : i+span])
		if prod > maxProd {
			maxProd = prod
		}
	}
	return maxProd, nil
}

func stringToDigits(s string) ([]int, error) {
	digits := make([]int, 0)
	for _, c := range s {
		digit, err := byteToDigit(c)
		if err != nil {
			return nil, err
		}
		digits = append(digits, digit)
	}
	return digits, nil
}

func byteToDigit(c rune) (int, error) {
	if c >= '0' && c <= '9' {
		return int(c) - int('0'), nil
	}
	return 0, errors.New("illegal digit character")
}

func mult(digits []int) int {
	prod := 1
	for _, digit := range digits {
		prod *= digit
	}
	return prod
}
