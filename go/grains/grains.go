package grains

import "errors"

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("illegal value")
	}
	return uint64(1) << uint(n-1), nil
}

func Total() uint64 {
	// The answer is Square(65) - 1, but that overflows uint64
	sq64, err := Square(64)
	if err == nil {
		return sq64 + (sq64 - 1)
	}
	return 0
}
