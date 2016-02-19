package brackets

const testVersion = 3

var openingBrackets = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}
var openings = []rune{'{', '(', '['}
var closings = []rune{'}', ')', ']'}

func Bracket(s string) (bool, error) {
	stack := make([]rune, len(s))
	stackLen := 0
	for _, c := range s {
		if isOpeningBacket(c) {
			stack[stackLen] = c
			stackLen++
		} else if isClosingBracket(c) {
			stackLen--
			if stackLen < 0 || stack[stackLen] != openingBrackets[c] {
				return false, nil
			}
		}
	}
	return stackLen == 0, nil
}

func isOpeningBacket(c rune) bool {
	for _, o := range openings {
		if o == c {
			return true
		}
	}
	return false
}

func isClosingBracket(c rune) bool {
	for _, o := range closings {
		if o == c {
			return true
		}
	}
	return false
}
