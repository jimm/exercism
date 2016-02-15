package triangle

import "math"

const TestVersion = 1

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

type Kind int

func KindFromSides(a, b, c float64) (k Kind) {
	switch {
	case isNotATriangle(a, b, c):
		k = NaT
	case isEquilateral(a, b, c):
		k = Equ
	case isIsoceles(a, b, c):
		k = Iso
	default:
		k = Sca
	}
	return
}

func isNotATriangle(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0 ||
		a+b < c || a+c < b || b+c < a ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
}

func isEquilateral(a, b, c float64) bool {
	return a == b && a == c
}

func isIsoceles(a, b, c float64) bool {
	return a == b || a == c || b == c
}
