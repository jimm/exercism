package brackets

const testVersion = 3

var pairs = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func Bracket(s string) (bool, error) {
	stack := make([]rune, len(s))
	stackLen := 0
	for _, c := range s {
		if isOpeningBracket(c) {
			stack[stackLen] = c
			stackLen++
		} else if stackLen > 0 && pairs[stack[stackLen-1]] == c {
			stackLen--
		} else {
			return false, nil
		}
	}
	return stackLen == 0, nil
}

func isOpeningBracket(c rune) bool {
	return pairs[c] != 0
}
