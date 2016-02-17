package cryptosquare

import (
	"strings"
	"regexp"
	"math"
)

const TestVersion = 1

func Encode(s string) string {
	strings.Join(ciphertextColumns(s), " ")
}

func normalizePlaintext(s string) string {
	re := regexp.MustCompile("\W+")
	return re.ReplaceAll(s, "")
}

func normalizeCiphertext(s string) string {
	return strings.Join(ciphertextColumns(s), " ")
}

func squareSize(s string) int {
	len := length(s)
	intSqrt = int(math.Sqrt(len))
	if intSqrt*intSqrt == len {
		return intSqrt
	}
	return intSqrt + 1
}

func plaintextSegments(s string) []string {
	segments = make([]string, 0)
	norm := normalizePlaintext(s)
	sqSize := squareSize(s)
	for i := 0; i < length(s); i += sqSize {
		segments = s[i:i+sqSize-1]
	}
	return segments
}

func ciphertextColumns(s string) []string {
	segs := plaintextSegments(normalizePlaintext(s))
	cols := length(segs[0])
	rows := length(segs)
	lastRowLen = length(segs[-1])
	retval := make([]string, rows)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			
			// (apply str (flatten
			//             (for [row (range rows)
			//                   :when (or (< row (dec rows)) (< col last-row-len))]
			//               (nth (nth segs row) col)))))))
		}
	}
	return retval
}
