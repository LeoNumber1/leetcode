package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},

		//{'C', 'A', 'A'},
		//{'A', 'A', 'A'},
		//{'B', 'C', 'D'},
	}

	word := "ABCCED"
	//word = "A"
	//word = "SEE"
	//word = "ABCB"

	//word = "AAB"
	//word = "DCB"
	//word = "AADB"
	//word = "BCD"

	//fmt.Println(exist(board, word))
	fmt.Println(exist1(board, word))
}

var direct = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右

//4 ms	3.6 MB
func exist(board [][]byte, word string) bool {
	//深度优先
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var dfs func(row, col, index int) bool

	dfs = func(row, col, index int) bool {
		if board[row][col] == word[index] {
			if index == len(word)-1 {
				return true
			}
			used[row][col] = true
			newRow, newColumn := 0, 0
			for _, v := range direct {
				newRow = row + v[0]
				newColumn = col + v[1]
				if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && !used[newRow][newColumn] {
					if dfs(newRow, newColumn, index+1) {
						return true
					}
				}
			}
		}
		used[row][col] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

//4 ms	3.3 MB
func exist1(board [][]byte, word string) bool {
	//深度优先
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	var dfs func(row, col, index int) bool

	dfs = func(row, col, index int) bool {
		tmp := board[row][col]
		if board[row][col] == word[index] {
			if index == len(word)-1 {
				return true
			}
			board[row][col] = ' '
			newRow, newColumn := 0, 0
			for _, v := range direct {
				newRow = row + v[0]
				newColumn = col + v[1]
				if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && board[newRow][newColumn] != ' ' {
					if dfs(newRow, newColumn, index+1) {
						return true
					}
				}
			}
		}
		board[row][col] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

//4 ms	3.7 MB
func existOfficial(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
