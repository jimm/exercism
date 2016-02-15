package pascal

import "fmt"

func Triangle(n int) (t [][]int) {
	fmt.Printf("x, n = %d\n", n)
	t = make([][]int, n)
	t[0] = make([]int, 1)
	fmt.Println("y")
	t[0][0] = 1
	fmt.Println("z")

	fmt.Println("a")
	for i := 1; i <= n-1; i++ {
		fmt.Println("b, i =", i)
		t[i] = make([]int, i)
		t[i][0] = 1
		t[i][i-1] = 1
		fmt.Println("b'")
		for j := 1; j < i-1; j++ {
			fmt.Println("  j", j)
			t[i][j] = t[i-1][j-1] + t[i-1][j]
		}
		fmt.Println("after b, t =", t)
	}
	fmt.Println("c, t =", t)
	return t
}
