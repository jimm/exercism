package prime

import "math"

const testVersion = 2

func Factors(n int64) []int64 {
	return factorsFor(n, 2, []int64{})
}

func factorsFor(n, prime int64, primeFactors []int64) []int64 {
	if n == 1 {
		return primeFactors
	}
	if n%prime == 0 {
		return factorsFor(n/prime, prime, append(primeFactors, prime))
	}
	return factorsFor(n, nextPrime(prime), primeFactors)
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
	ceiling := int64(math.Sqrt(float64(n)))
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
	i++
	return primeTest(n, i, j, h)
}
