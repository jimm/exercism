package prime

import "math"

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	p := int64(2)
	for i := 1; i < n; i++ {
		p = nextPrime(p)
	}
	return int(p), true
}

func nextPrime(n int64) int64 {
	if n == 2 {
		return 3
	}
	for n = n + 2; !isPrime(n); n += 2 {
	}
	return n
}

func isPrime(n int64) bool {
	ceiling := round(math.Sqrt(float64(n)))
	if n == ceiling*ceiling {
		return false
	}
	return primeTest(n, 2, ceiling, ceiling)
}

func primeTest(n, i, j, h int64) bool {
	if i == j {
		return i == h
	}
	if n%i == 0 {
		return i == h
	}
	return primeTest(n, i+1, j, h)
}

// Is there a built-in round() somewhere? There's no math.Round().
func round(f float64) int64 {
	i64 := int64(f)
	if i64 == int64(f+0.5) {
		return i64
	}
	return i64 + 1
}
