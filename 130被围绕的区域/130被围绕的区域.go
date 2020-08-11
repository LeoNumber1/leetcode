package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X'},
	}

	//solve(board)
	solveOfficial2(board)

	fmt.Println(board)
}

var count = [][]int{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	matrix := make([][]int, m) //定义一个同等大小的二维数组，值为0——初始值；为1——值不用动；为2——值不用动，且是边界'O'
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 || matrix[i][j] == 2 {
				continue
			}

			if board[i][j] == 'X' {
				matrix[i][j] = 1 //1代表不用改动
				continue
			}

			//判断是边上，并且值为'O'的情况
			if (i == 0 || i == m-1) || (j == 0 || j == n-1) {
				matrix[i][j] = 2 //值不用动，且是边界0
				dfs(board, matrix, i, j)
			}
		}
	}
	//fmt.Println("after :", matrix)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && matrix[i][j] == 0 {
				board[i][j] = 'X'
			}
		}
	}
}

func solve1(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	matrix := make([][]int, m) //定义一个同等大小的二维数组，值为0——初始值；为1——值不用动；为2——值不用动，且是边界'O'
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	queue := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 || matrix[i][j] == 2 {
				continue
			}

			if board[i][j] == 'X' {
				matrix[i][j] = 1 //1代表不用改动
				continue
			}

			//判断是边上，并且值为'O'的情况
			if (i == 0 || i == m-1) || (j == 0 || j == n-1) {
				matrix[i][j] = 2 //值不用动，且是边界0
				dfs(board, matrix, i, j)
			} else {
				queue = append(queue, []int{i, j})
			}
		}
	}
	//fmt.Println("after :", matrix)
	//
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if board[i][j] == 'O' && matrix[i][j] == 0 {
	//			board[i][j] = 'X'
	//		}
	//	}
	//}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		if matrix[point[0]][point[1]] == 0 {
			board[point[0]][point[1]] = 'X'
		}
	}
}

func dfs(board [][]byte, matrix [][]int, x, y int) {
	for i := 0; i < 4; i++ { //循环4次情况
		newRow := x + count[i][0]
		newColumn := y + count[i][1]

		if newRow >= 0 && newRow < len(board) && newColumn >= 0 && newColumn < len(board[0]) { //新值在范围内
			if matrix[newRow][newColumn] == 1 || matrix[newRow][newColumn] == 2 {
				continue
			}

			if board[newRow][newColumn] == 'X' {
				matrix[newRow][newColumn] = 1 //1代表不用改动
				continue
			}

			matrix[newRow][newColumn] = 2 //值不用动，且是边界0
			dfs(board, matrix, newRow, newColumn)
		}
	}
}

var n1, m1 int

//深度优先搜索
func solveOfficial(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n1, m1 = len(board), len(board[0])
	//判断所有行边界O和其相连的O，把它们全改为A
	for i := 0; i < n1; i++ {
		dfs1(board, i, 0)
		dfs1(board, i, m1-1)
	}
	//判断所有列边界O和其相连的O，把它们全改为A
	for i := 1; i < m1-1; i++ {
		dfs1(board, 0, i)
		dfs1(board, n1-1, i)
	}
	//循环原始数组，把A的全改回O，O的全改为X
	for i := 0; i < n1; i++ {
		for j := 0; j < m1; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs1(board [][]byte, x, y int) {
	if x < 0 || x >= n1 || y < 0 || y >= m1 || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs1(board, x+1, y)
	dfs1(board, x-1, y)
	dfs1(board, x, y+1)
	dfs1(board, x, y-1)
}

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)

//广度优先搜索
func solveOfficial2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	queue := [][]int{}
	//如果是行边界O，则加入队列queue
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}
		if board[i][m-1] == 'O' {
			queue = append(queue, []int{i, m - 1})
		}
	}
	//如果是列边界O，则加入队列queue
	for i := 1; i < m-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
		}
		if board[n-1][i] == 'O' {
			queue = append(queue, []int{n - 1, i})
		}
	}
	//广度优先，将每个边界放入队列，在队列里找和它相连的O，改为A后，再放入队列
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx < 0 || my < 0 || mx >= n || my >= m || board[mx][my] != 'O' {
				continue
			}
			queue = append(queue, []int{mx, my})
		}
	}
	//同上，将数组还原
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
