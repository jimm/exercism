package diffsquares

func SquareOfSums(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	nSquared := n * n
	nCubed := nSquared * n
	return int(float64(nCubed)/3.0 + float64(nSquared)/2.0 + float64(n)/6.0)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
