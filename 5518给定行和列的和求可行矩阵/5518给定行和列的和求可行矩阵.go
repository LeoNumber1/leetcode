package main

import "fmt"

func main() {
	rowSum := []int{3, 8}
	colSum := []int{4, 7}

	rowSum = []int{5, 7, 10}
	colSum = []int{8, 6, 8}

	fmt.Println(restoreMatrix(rowSum, colSum))
}

//	84 ms	8.7 MB
func restoreMatrix1(rowSum []int, colSum []int) [][]int {
	m := len(rowSum)
	n := len(colSum)
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			minVal := min(rowSum[i], colSum[j])
			arr[i][j] = minVal
			rowSum[i] -= minVal
			colSum[j] -= minVal
			//剪枝：当发现rowSum和colSum都是0的时候，就没必要再继续遍历下去了
			if rowSum[i] == 0 && colSum[j] == 0 {
				break
			}
		}
	}
	return arr
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m := len(rowSum)
	n := len(colSum)
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			minVal := min(rowSum[i], colSum[j])
			arr[i][j] = minVal
			rowSum[i] -= minVal
			colSum[j] -= minVal
			//剪枝：当发现rowSum和colSum都是0的时候，就没必要再继续遍历下去了
			if rowSum[i] == 0 && colSum[j] == 0 {
				break
			}
		}
	}
	return arr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
