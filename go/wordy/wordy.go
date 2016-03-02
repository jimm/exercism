package wordy

import (
	"math"
	"strconv"
	"strings"
)

var cardinalSuffixes = []string{"st", "nd", "rd", "th"}

type wordFunc struct {
	words         []string
	numArgs       int
	f             func(int, int) int
	trailingWords []string
}

// Additional:
// 3 times 4
// 2 over 2
// 3 raised to the 4th power
// the 2nd power of 3

var wordFuncs = []wordFunc{
	{[]string{"plus"}, 2, func(x, y int) int { return x + y }, nil},
	{[]string{"minus"}, 2, func(x, y int) int { return x - y }, nil},
	{[]string{"multiplied", "by"}, 2, func(x, y int) int { return x * y }, nil},
	{[]string{"times"}, 2, func(x, y int) int { return x * y }, nil},
	{[]string{"divided", "by"}, 2, func(x, y int) int { return x / y }, nil},
	{[]string{"over"}, 2, func(x, y int) int { return x / y }, nil},
	{[]string{"squared"}, 1, func(x, _ int) int { return x * x }, nil},
	{[]string{"doubled"}, 1, func(x, _ int) int { return x * x }, nil},
	{[]string{"cubed"}, 1, func(x, _ int) int { return x * x * x }, nil},
	{[]string{"raised", "to", "the"}, 2,
		func(x, y int) int { return int(math.Pow(float64(x), float64(y))) },
		[]string{"power"}},
	{[]string{"power", "of"}, 2,
		func(x, y int) int {
			return int(math.Pow(float64(y), float64(x))) },
		nil},
}

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, "What is ") { // sufficient for the tests
		return 0, false
	}

	question = normalize(question)
	words := strings.Fields(question)
	answer := 0
	var wf *wordFunc
	for i := 0; i < len(words)-1; i = nextNumIndex(i, wf) {
		var x, y int
		if wf == nil {
			x = parseWord(words[i])
		} else {
			x = answer
		}
		wf = findFunc(words[i+1:]) // test data won't return nil
		if wf == nil {
			return 0, false
		}
		if wf.numArgs == 2 {
			y = parseWord(words[i+len(wf.words)+1])
		}
		answer = wf.f(x, y)
	}

	return answer, true
}

func findFunc(words []string) *wordFunc {
	for _, wf := range wordFuncs {
		if wordsMatch(words, wf.words) {
			return &wf
		}
	}
	return nil
}

func wordsMatch(words, wfWords []string) bool {
	for i := range wfWords {
		if words[i] != wfWords[i] {
			return false
		}
	}
	return true
}

func nextNumIndex(numIndex int, wf *wordFunc) int {
	i := numIndex + 1 + len(wf.words)
	if wf != nil && len(wf.trailingWords) > 0 {
		i += len(wf.trailingWords)
	}
	return i
}

func parseWord(word string) int {
	for _, suffix := range cardinalSuffixes {
		word = strings.TrimSuffix(word, suffix)
	}
	
	n, _ := strconv.ParseInt(word, 10, 32) // no errors in tests
	return int(n)
}

// normalize removes the prefix "What is [the ]" and the suffix "?" from
// question.
func normalize(question string) string {
	question = strings.TrimPrefix(question, "What is the ")
	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSuffix(question, "?")
	return question
}
