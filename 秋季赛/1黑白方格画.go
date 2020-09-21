package main

import "fmt"

func main() {
	n := 2
	k := 2

	fmt.Println(paintingPlan(n, k))
}

func paintingPlan(n int, k int) int {
	if k < n {
		return 0
	}
	if k == n*n {
		return 1
	}
	if k == n {
		return n * 2
	}
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}
	var ans, sheng int = 0, k
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count := 0
			if !grid[i][j] {
				grid[i][j] = true
				count++
			}
			sheng -= count
			if sheng == 0 {
				ans++
			}
		}
	}
}
