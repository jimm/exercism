package summultiples

func MultipleSummer(xs ...int) func(int) int {
	return func(limit int) int {
		sum := 0
		for i := 0; i < limit; i++ {
			if isMultipleOfAny(i, xs) {
				sum += i
			}
		}
		return sum
	}
}

func isMultipleOfAny(i int, xs []int) bool {
	for _, x := range xs {
		if i%x == 0 {
			return true
		}
	}
	return false
}
