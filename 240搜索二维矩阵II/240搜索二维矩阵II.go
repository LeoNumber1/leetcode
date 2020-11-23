package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5

	target = 20

	fmt.Println(searchMatrixOfficial2(matrix, target))
}

//28 ms	6.6 MB
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		v := matrix[row][col]
		if target == v {
			return true
		} else if target > v {
			col++
		} else {
			row--
		}
	}
	return false
}

//28 ms-85.77%	6.6 MB-45.38%	二分法
func searchMatrixOfficial2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	binarySearch := func(start int, vertical bool) bool {
		lo, hi := start, n-1
		if vertical {
			hi = m - 1
		}
		for hi >= lo {
			//mid := (lo + hi)/2
			mid := lo + (hi-lo)/2
			if vertical { //列内搜索
				if matrix[mid][start] < target {
					lo = mid + 1
				} else if matrix[mid][start] > target {
					hi = mid - 1
				} else {
					return true
				}
			} else { //行内搜索
				if matrix[start][mid] < target {
					lo = mid + 1
				} else if matrix[start][mid] > target {
					hi = mid - 1
				} else {
					return true
				}
			}
		}
		return false
	}
	shorterDim := 0
	if m > n {
		shorterDim = n
	} else {
		shorterDim = m
	}
	for i := 0; i < shorterDim; i++ {
		if binarySearch(i, true) || binarySearch(i, false) {
			return true
		}
	}
	return false
}
