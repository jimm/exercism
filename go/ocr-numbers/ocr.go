package ocr

import "strings"

type segmentsToDigit struct {
	segments []string
	digit    string
}

var translator = []segmentsToDigit{
	{[]string{" _ ", "| |", "|_|"}, "0"},
	{[]string{"   ", "  |", "  |"}, "1"},
	{[]string{" _ ", " _|", "|_ "}, "2"},
	{[]string{" _ ", " _|", " _|"}, "3"},
	{[]string{"   ", "|_|", "  |"}, "4"},
	{[]string{" _ ", "|_ ", " _|"}, "5"},
	{[]string{" _ ", "|_ ", "|_|"}, "6"},
	{[]string{" _ ", "  |", "  |"}, "7"},
	{[]string{" _ ", "|_|", "|_|"}, "8"},
	{[]string{" _ ", "|_|", " _|"}, "9"},
}

func Recognize(allLines string) []string {
	recognized := []string{}
	lines := strings.Split(allLines, "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			digits := ""
			for _, ds := range segmentsFromLines(lines[i : i+3]) {
				digits += recognizeDigit(ds)
			}
			recognized = append(recognized, digits)
			i += 3 // we ate three lines, plus one blank
		}
	}
	return recognized
}

func recognizeDigit(segments []string) string {
	for _, s2d := range translator {
		if segmentsMatch(s2d.segments, segments) {
			return s2d.digit
		}
	}
	return "?"
}

func segmentsFromLines(lines []string) [][]string {
	numDigits := len(lines[0]) / 3
	digitSegments := make([][]string, numDigits)
	segmentIndex := 0
	for col := 0; col < len(lines[0]); col++ {
		digitSegments[segmentIndex] = []string{}
		for row := 0; row < 3; row++ {
			digitSegments[segmentIndex] = append(digitSegments[segmentIndex], lines[row][col:col+3])
		}
		col += 2 // we ate three chars
		segmentIndex++
	}
	return digitSegments
}

func segmentsMatch(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
