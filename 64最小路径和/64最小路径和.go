package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			if j == n-1 && i != m-1 {
				grid[i][j] += grid[i+1][j]
			} else if j != n-1 && i == m-1 {
				grid[i][j] += grid[i][j+1]
			} else {
				grid[i][j] += min(grid[i][j+1], grid[i+1][j])
			}
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
