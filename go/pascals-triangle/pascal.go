package pascal

func Triangle(n int) (t [][]int) {
	t = initTriangle(n)
	for row := 2; row < n; row++ {
		for j := 1; j < row; j++ {
			t[row][j] = t[row-1][j-1] + t[row-1][j]
		}
	}
	return
}

// initTriangle allocates memory and initializes the outer 1's.
func initTriangle(n int) (t [][]int) {
	t = make([][]int, n)
	for row := 0; row < n; row++ {
		t[row] = make([]int, row+1)
		t[row][0] = 1
		t[row][row] = 1
	}
	return
}
