package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const TestVersion = 1

func Encode(s string) string {
	return strings.Join(ciphertextColumns(s), " ")
}

func normalizePlaintext(s string) string {
	re := regexp.MustCompile("\\W+")
	return strings.ToLower(string(re.ReplaceAll([]byte(s), []byte(""))))
}

func normalizeCiphertext(s string) string {
	return strings.Join(ciphertextColumns(s), " ")
}

func squareSize(s string) int {
	len := len(s)
	intSqrt := int(math.Sqrt(float64(len)))
	if intSqrt*intSqrt == len {
		return intSqrt
	}
	return intSqrt + 1
}

func plaintextSegments(s string) []string {
	segments := make([]string, 0)
	norm := normalizePlaintext(s)
	sqSize := squareSize(norm)
	for i := 0; i < len(s); i += sqSize - 1 {
		segments = append(segments, s[i:i+sqSize-1])
	}
	return segments
}

func ciphertextColumns(s string) []string {
	norm := normalizePlaintext(s)
	if len(norm) == 0 {
		return make([]string, 0)
	}

	sqSize := squareSize(norm)
	retval := make([]string, sqSize)

	col := -1
	row := 0
	for i := 0; i < len(norm); i++ {
		col += 1
		if col == sqSize {
			col = 0
			row += 1
		}
		retval[col] += norm[i : i+1]
	}
	return retval
}
