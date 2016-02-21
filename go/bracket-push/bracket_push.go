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
		switch {
		case isOpeningBracket(c):
			stack[stackLen] = c
			stackLen++
		case stackLen > 0 && pairs[stack[stackLen-1]] == c:
			stackLen--
		default:
			return false, nil
		}
	}
	return stackLen == 0, nil
}

func isOpeningBracket(c rune) bool {
	for k := range pairs {
		if c == k {
			return true
		}
	}
	return false
}
