package slice

func All(n int, s string) []string {
	ss := make([]string, 0)
	for i := 0; i < len(s)-n+1; i++ {
		ss = append(ss, s[i:n+i])
	}
	return ss
}

func Frist(n int, s string) string {
	if n >= len(s) {
		n = len(s)
	}
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		ok = false
	} else {
		first = s[:n]
		ok = true
	}
	return
}
