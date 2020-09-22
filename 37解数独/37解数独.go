package main

import (
	"fmt"
	"math/bits"
)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board = [][]byte{
		{'.', '.', '.', '2', '.', '.', '.', '.', '3'},
		{'.', '.', '2', '.', '8', '.', '.', '9', '.'},
		{'.', '7', '.', '.', '.', '9', '.', '.', '.'},
		{'3', '.', '.', '8', '.', '.', '.', '.', '2'},
		{'.', '.', '.', '.', '.', '7', '.', '.', '5'},
		{'.', '9', '8', '.', '.', '1', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '8', '2', '.'},
		{'2', '.', '.', '3', '6', '.', '.', '.', '.'},
		{'6', '.', '.', '.', '.', '.', '7', '.', '.'},
	}

	//solveSudoku(board)
	solveSudokuOfficial1(board)
	//fmt.Println(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
	_ = board

	//for i := 0; i < 9; i++ {
	//	for j := 0; j < 9; j++ {
	//		fmt.Print(getGrid(i, j))
	//	}
	//	fmt.Println()
	//}

	fmt.Println()
}

//36 ms	2.2 MB
func solveSudoku(board [][]byte) {
	row := map[int]map[byte]bool{}
	column := map[int]map[byte]bool{}
	grid := map[int]map[byte]bool{}

	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		column[i] = make(map[byte]bool)
		grid[i] = make(map[byte]bool)
	}

	var spaces [][2]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row[i][board[i][j]] = true
				column[j][board[i][j]] = true
				grid[getGrid(i, j)][board[i][j]] = true
			} else {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		index := getGrid(i, j)
		for num := 1; num <= 9; num++ {

			b := byte(48 + num)
			if _, has := row[i][b]; has {
				continue
			}
			if _, has := column[j][b]; has {
				continue
			}
			if _, has := grid[index][b]; has {
				continue
			}

			board[i][j] = b
			row[i][b], column[j][b], grid[index][b] = true, true, true
			if dfs(pos + 1) {
				return true
			}
			//board[i][j] = '.'
			delete(row[i], b)
			delete(column[j], b)
			delete(grid[index], b)
		}

		return false
	}
	dfs(0)
	return
}

func getGrid(i, j int) int {
	x := i / 3
	y := j / 3
	switch i {
	case 3, 4, 5:
		y += 2
	case 6, 7, 8:
		y += 4
	}
	return x + y
}

//0 ms	2.1 MB
func solveSudokuOfficial(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

//0 ms	2.1 MB	位运算
func solveSudokuOfficial1(board [][]byte) {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int

	flip := func(i, j int, digit byte) {
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				flip(i, j, digit)
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3]) // 0x1ff 即二进制的 9 个 1
		for ; mask > 0; mask &= mask - 1 {                       // 最右侧的 1 置为 0
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
}

//0 ms	2.1 MB 位运算+确定值
func solveSudokuOfficial2(board [][]byte) {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int

	flip := func(i, j int, digit byte) {
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}

	for i, row := range board {
		for j, b := range row {
			if b != '.' {
				digit := b - '1'
				flip(i, j, digit)
			}
		}
	}

	for {
		modified := false
		for i, row := range board {
			for j, b := range row {
				if b != '.' {
					continue
				}
				mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3])
				if mask&(mask-1) == 0 { // mask 的二进制表示仅有一个 1
					digit := byte(bits.TrailingZeros(mask))
					flip(i, j, digit)
					board[i][j] = digit + '1'
					modified = true
				}
			}
		}
		if !modified {
			break
		}
	}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3]) // 0x1ff 即二进制的 9 个 1
		for ; mask > 0; mask &= mask - 1 {                       // 最右侧的 1 置为 0
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
}
