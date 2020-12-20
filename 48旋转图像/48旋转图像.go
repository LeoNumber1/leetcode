package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	//matrix = [][]int{
	//	{5, 1, 9, 11},
	//	{2, 4, 8, 10},
	//	{13, 3, 6, 7},
	//	{15, 14, 12, 16},
	//}

	//rotate(matrix)
	rotateOfficial2(matrix)

	fmt.Println(matrix)
}

//0 ms-100.00%	2.2 MB-94.80%
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
	//step1:上下翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	//step2:主对角线翻转
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

//0 ms-100.00%	2.3 MB-5.39%
func rotateOfficial1(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
}

//0 ms-100.00%	2.2 MB-94.80%
func rotateOfficial2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
