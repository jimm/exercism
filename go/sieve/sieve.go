package sieve

func Sieve(n int) (primes []int) {
	primes = make([]int, 0)

	// Cheat by pre-filling our list of ints with 2 and start with only odd
	// ints. It would be more efficient to pre-allocate the whole slice
	// instead of using append. However, then I'd have to use a separate
	// index variable to remember where we are in the slice when filling it.
	// This is simpler and good enough for small values of n.
	ints := make([]int, 1)
	ints[0] = 2
	for i := 3; i <= n; i += 2 {
		ints = append(ints, i)
	}

	for len(ints) > 0 {
		primes = append(primes, ints[0])
		ints = Filter(ints[1:], ints[0])
	}

	return primes
}

func Filter(ns []int, p int) []int {
	var filtered []int
	for _, i := range ns {
		if i%p != 0 {
			filtered = append(filtered, i)
		}
	}
	return filtered
}
