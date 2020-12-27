package main

import "fmt"

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	//matrix = [][]byte{}
	//matrix = [][]byte{{'1'}}
	//matrix = [][]byte{{'0', '1'}}
	//matrix = [][]byte{{'1', '0'}, {'1', '0'}}
	//matrix = [][]byte{{'0', '1'}, {'0', '1'}}

	//fmt.Println(maximalRectangle(matrix))
	fmt.Println(maximalRectangleOfficial(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	var ans int
	left := make([][]int, m) //定义left[i][j],值为matrix[i][j]左边1的个数
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					left[i][j] = 1
				} else {
					left[i][j] = left[i][j-1] + 1
				}
				if left[i][j] > ans {
					ans = left[i][j]
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if left[i][j] != 0 {
				width := left[i][j]

				height := 0
				for k := i; k >= 0; k-- {
					if left[k][j] == 0 {
						break
					}
					height++
					width = min(width, left[k][j])
					temp := width * height
					if temp > ans {
						ans = temp
					}
				}
			}
		}
	}

	return ans
}

func maximalRectangleOfficial(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
