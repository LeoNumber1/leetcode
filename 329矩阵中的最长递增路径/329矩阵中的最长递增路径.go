package main

import "fmt"

func main() {
	matrix := [][]int{
		//{9, 9, 4},
		//{6, 6, 8},
		//{2, 1, 1},
		//{3, 4, 5},
		//{3, 2, 6},
		//{2, 2, 1},
		//{7, 6, 1, 1},
		//{2, 7, 6, 0},
		//{1, 3, 5, 1},
		//{6, 6, 3, 2},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
		{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
		{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
		{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(longestIncreasingPath2(matrix))
	fmt.Println(longestIncreasingPathOfficial1(matrix))
	fmt.Println(longestIncreasingPathOfficial2(matrix))
}

var m, n int
var matrixPublic [][]int
var num = [][]int{{}}

func longestIncreasingPath(matrix [][]int) int {
	m = len(matrix)
	if m == 0 {
		return 0
	}
	n = len(matrix[0])
	matrixPublic = matrix
	num = make([][]int, m)
	for i := 0; i < m; i++ {
		num[i] = make([]int, n)
	}

	var ret int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num[i][j] = getLen(i, j) + 1
			if i == 0 && j == 0 {
				ret = num[i][j]
			} else {
				if ret < num[i][j] {
					ret = num[i][j]
				}
			}
		}
	}
	fmt.Println(num)
	return ret
}

func getLen(i, j int) int {
	var lenthUp, lenthDown, lenthLeft, lenthRight int
	if i+1 < m && matrixPublic[i+1][j] > matrixPublic[i][j] {
		i += 1
		lenthDown = getLen(i, j) + 1
		i -= 1
		//num[i][j] += 1
	}

	if j+1 < n && matrixPublic[i][j+1] > matrixPublic[i][j] {
		j += 1
		lenthRight = getLen(i, j) + 1
		j -= 1
	}

	if i > 0 && matrixPublic[i-1][j] > matrixPublic[i][j] {
		i -= 1
		lenthUp = getLen(i, j) + 1
		i += 1
		//lenthUp = num[i-1][j]
	}

	if j > 0 && matrixPublic[i][j-1] > matrixPublic[i][j] {
		j -= 1
		lenthLeft = getLen(i, j) + 1
		j += 1
		//lenthLeft = num[i][j-1]
	}
	return getMax(lenthUp, lenthDown, lenthLeft, lenthRight)
}

func getMax(lenthUp, lenthDown, lenthLeft, lenthRight int) int {
	var max1, max2 int
	if lenthUp > lenthDown {
		max1 = lenthUp
	} else {
		max1 = lenthDown
	}

	if lenthLeft > lenthRight {
		max2 = lenthLeft
	} else {
		max2 = lenthRight
	}

	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}

func longestIncreasingPath2(matrix [][]int) int {
	m = len(matrix)
	if m == 0 {
		return 0
	}
	n = len(matrix[0])
	matrixPublic = matrix
	num = make([][]int, m)
	for i := 0; i < m; i++ {
		num[i] = make([]int, n)
	}

	var ret int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if num[i][j] == 0 {
				num[i][j] = getLen2(i, j) + 1
			}
			if i == 0 && j == 0 {
				ret = num[i][j]
			} else {
				if ret < num[i][j] {
					ret = num[i][j]
				}
			}
		}
	}
	//fmt.Println(num)
	return ret
}

func getLen2(i, j int) int {
	var lenthUp, lenthDown, lenthLeft, lenthRight int
	if i+1 < m && matrixPublic[i+1][j] > matrixPublic[i][j] {
		i += 1
		if num[i][j] == 0 {
			lenthDown = getLen2(i, j) + 1
		} else {
			lenthDown = num[i][j]
		}
		num[i][j] = lenthDown
		i -= 1
	}

	if j+1 < n && matrixPublic[i][j+1] > matrixPublic[i][j] {
		j += 1
		if num[i][j] == 0 {
			lenthRight = getLen2(i, j) + 1
		} else {
			lenthRight = num[i][j]
		}
		num[i][j] = lenthRight
		j -= 1
	}

	if i > 0 && matrixPublic[i-1][j] > matrixPublic[i][j] {
		i -= 1
		if num[i][j] == 0 {
			lenthUp = getLen2(i, j) + 1
		} else {
			lenthUp = num[i][j]
		}
		num[i][j] = lenthUp
		i += 1
	}

	if j > 0 && matrixPublic[i][j-1] > matrixPublic[i][j] {
		j -= 1
		if num[i][j] == 0 {
			lenthLeft = getLen2(i, j) + 1
		} else {
			lenthLeft = num[i][j]
		}
		num[i][j] = lenthLeft
		j += 1
	}
	return getMax(lenthUp, lenthDown, lenthLeft, lenthRight)
}

var (
	dirs          = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	rows, columns int
)

func longestIncreasingPathOfficial1(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, memo))
		}
	}
	//fmt.Println(memo)
	return ans
}

func dfs(matrix [][]int, row, column int, memo [][]int) int {
	if memo[row][column] != 0 {
		return memo[row][column]
	}
	memo[row][column]++
	for _, dir := range dirs {
		newRow, newColumn := row+dir[0], column+dir[1]
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
			memo[row][column] = max(memo[row][column], dfs(matrix, newRow, newColumn, memo)+1)
		}
	}
	return memo[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestIncreasingPathOfficial2(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	outdegrees := make([][]int, rows)
	for i := 0; i < rows; i++ {
		outdegrees[i] = make([]int, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for _, dir := range dirs {
				newRow, newColumn := i+dir[0], j+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[i][j] {
					outdegrees[i][j]++
				}
			}
		}
	}

	queue := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if outdegrees[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	ans := 0
	for len(queue) != 0 {
		ans++
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]
			row, column := cell[0], cell[1]
			for _, dir := range dirs {
				newRow, newColumn := row+dir[0], column+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] < matrix[row][column] {
					outdegrees[newRow][newColumn]--
					if outdegrees[newRow][newColumn] == 0 {
						queue = append(queue, []int{newRow, newColumn})
					}
				}
			}
		}
	}
	return ans
}
