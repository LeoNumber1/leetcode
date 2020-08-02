package main

import (
	"fmt"
	"strconv"
)

func main() {
	maze := []string{
		"S#O", "M..", "M.T",
	}
	//fmt.Println(minimalSteps(maze))
	//fmt.Println(mazeGlobal)
	fmt.Println(minimalStepsOfficial(maze))
}

var mazeGlobal [][]int
var moves = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}
var rows, columns int

func minimalSteps(maze []string) int {
	S, T := make([]int, 2), make([]int, 2)
	var M = make([][]int, 0)
	var O = make([][]int, 0)
	var road = make(map[string]int)
	rows, columns = len(maze), len(maze[0])
	mazeGlobal = make([][]int, rows)
	for i := 0; i < rows; i++ {
		mazeGlobal[i] = make([]int, columns)
	}

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			slice := make([]int, 2)
			slice[0] = i
			slice[1] = j
			switch string(maze[i][j]) {
			case "S":
				mazeGlobal[i][j] = 1
				S[0] = i
				S[1] = j
				road[strconv.Itoa(i)+","+strconv.Itoa(j)] = 1
			case "T":
				mazeGlobal[i][j] = 1
				T[0] = i
				T[1] = j
				road[strconv.Itoa(i)+","+strconv.Itoa(j)] = 1
			case "O":
				mazeGlobal[i][j] = 1
				O = append(O, slice)
				road[strconv.Itoa(i)+","+strconv.Itoa(j)] = 1
			case "M":
				mazeGlobal[i][j] = 1
				M = append(M, slice)
				road[strconv.Itoa(i)+","+strconv.Itoa(j)] = 1
			case ".":
				mazeGlobal[i][j] = 1
				road[strconv.Itoa(i)+","+strconv.Itoa(j)] = 1
			case "#":
				mazeGlobal[i][j] = 0
			}
		}
	}
	//fmt.Println(S, M, T, O, road)
	mLen := len(M)
	oLen := len(O)
	if mLen > 0 && oLen == 0 { //有机关没石头
		return -1
	}

	var start = S

	var target [][]int

	if oLen > 0 && mLen > 0 { //有石头有机关
		//minRoadTarget(start, []int{1, 1})
		minRoadTarget([]int{1, 1}, start)
		//for i:=0;i<len() {
		//
		//}
		//
		//target = O[0]
	} else { //没有机关
		target = append(target, T)
	}

	for i := start[0]; i < len(maze); i++ {
		for j := start[1]; j < len(maze[i]); j++ {

		}
	}

	return -1
}

func minRoadTarget(start []int, target []int) int {
	//for i := start[0]; i < len(mazeGlobal); i++ {
	//
	//}
	var maze [][]int
	maze = make([][]int, rows)
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		maze[i] = make([]int, columns)
		memo[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			memo[i][j] = rows * columns
		}
		copy(maze[i], mazeGlobal[i])
	}
	//maze[0][0] = 5
	//fmt.Println("in func maze:", maze)
	var ans int = rows * columns
	//
	//for i:=target[0];i>=0;i-- {
	//
	//}

	defRet := def1(maze, start[0], start[1], start, target)
	if defRet == 0 {
		return -1
	}
	ans = min(ans, defRet) - 1

	//maze[start[0]][start[1]] = 0
	//for _, move := range moves {
	//	newRow, newColumn := start[0]+move[0], start[1]+move[1]
	//	if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && maze[newRow][newColumn] == 1 {
	//ans = min(ans, dfs(maze, start[0], start[1], target, memo))
	//	}
	//}
	return ans
}

func def1(maze [][]int, row, column int, start, target []int) int {
	if maze[row][column] != 1 && maze[row][column] != 0 {
		return maze[row][column]
	}
	for _, move := range moves {
		newRow, newColumn := row+move[0], column+move[1]
		if newRow == target[0] && newColumn == target[1] {
			maze[newRow][newColumn] += maze[row][column]
			return maze[newRow][newColumn]
		}
		if newRow == start[0] && newColumn == start[1] {
			continue
		}
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && maze[newRow][newColumn] == 1 {
			maze[newRow][newColumn] += maze[row][column]
		}
	}

	for _, move := range moves {
		newRow, newColumn := row+move[0], column+move[1]
		//if newRow == target[0] && newColumn == target[1] {
		//
		//}
		if newRow == start[0] && newColumn == start[1] {
			continue
		}
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && maze[newRow][newColumn] != 1 && maze[newRow][newColumn] != 0 {
			//maze[newRow][newColumn] += maze[row][column]
			def1(maze, newRow, newColumn, start, target)
		}
	}
	return 0
}

//func dfs(maze [][]int, row, column int, target []int, memo [][]int) int {
//	if memo[row][column] != rows*columns {
//		return memo[row][column]
//	} else {
//		memo[row][column] = 0
//	}
//	memo[row][column]++
//	//maze[row][column] = 0
//	for _, move := range moves {
//		newRow, newColumn := row+move[0], column+move[1]
//		if newRow == target[0] && newColumn == target[1] {
//			break
//		}
//		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && maze[newRow][newColumn] == 1 {
//			memo[row][column] = min(memo[row][column], dfs(maze, newRow, newColumn, target, memo)+1)
//		}
//	}
//	return memo[row][column]
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
	以上做了一半放弃了，下面是官方解答
*/
var (
	dx   = []int{1, -1, 0, 0}
	dy   = []int{0, 0, 1, -1}
	n, m int
)

func minimalStepsOfficial(maze []string) int {
	n, m = len(maze), len(maze[0])
	// 机关 & 石头
	var buttons, stones [][]int
	// 起点 & 终点
	sx, sy, tx, ty := -1, -1, -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch maze[i][j] {
			case 'M':
				buttons = append(buttons, []int{i, j})
			case 'O':
				stones = append(stones, []int{i, j})
			case 'S':
				sx, sy = i, j
			case 'T':
				tx, ty = i, j
			}
		}
	}

	numBlock, numStone := len(buttons), len(stones)
	startDist := bfs(sx, sy, maze)
	// 边界情况：没有机关
	if numBlock == 0 {
		return startDist[tx][ty]
	}
	// 从某个机关到其他机关 / 起点与终点的最短距离。
	dist := make([][]int, numBlock)
	for i := 0; i < numBlock; i++ {
		dist[i] = make([]int, numBlock+2)
		for j := 0; j < numBlock+2; j++ {
			dist[i][j] = -1
		}
	}
	// 中间结果
	dd := make([][][]int, numBlock)
	for i := 0; i < numBlock; i++ {
		dd[i] = bfs(buttons[i][0], buttons[i][1], maze) //计算地图上任意一点到机关的距离
		// 从某个点到终点不需要拿石头
		dist[i][numBlock+1] = dd[i][tx][ty]
	}
	for i := 0; i < numBlock; i++ {
		tmp := -1
		for k := 0; k < numStone; k++ {
			midX, midY := stones[k][0], stones[k][1]
			if dd[i][midX][midY] != -1 && startDist[midX][midY] != -1 {
				if tmp == -1 || tmp > dd[i][midX][midY]+startDist[midX][midY] {
					tmp = dd[i][midX][midY] + startDist[midX][midY]
				}
			}
		}
		dist[i][numBlock] = tmp
		for j := i + 1; j < numBlock; j++ {
			mn := -1
			for k := 0; k < numStone; k++ {
				midX, midY := stones[k][0], stones[k][1]
				if dd[i][midX][midY] != -1 && startDist[midX][midY] != -1 {
					if mn == -1 || mn > dd[i][midX][midY]+dd[j][midX][midY] {
						mn = dd[i][midX][midY] + dd[j][midX][midY]
					}
				}
			}
			dist[i][j] = mn
			dist[j][i] = mn
		}
	}
	// 无法达成的情形
	for i := 0; i < numBlock; i++ {
		if dist[i][numBlock] == -1 || dist[i][numBlock+1] == -1 {
			return -1
		}
	}
	// dp 数组， -1 代表没有遍历到
	dp := make([][]int, 1<<numBlock)
	for i := 0; i < (1 << numBlock); i++ {
		dp[i] = make([]int, numBlock)
		for j := 0; j < numBlock; j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < numBlock; i++ {
		dp[1<<i][i] = dist[i][numBlock]
	}

	// 由于更新的状态都比未更新的大，所以直接从小到大遍历即可
	for mask := 1; mask < (1 << numBlock); mask++ {
		for i := 0; i < numBlock; i++ {
			// 当前 dp 是合法的
			if mask&(1<<i) != 0 {
				for j := 0; j < numBlock; j++ {
					// j 不在 mask 里
					if mask&(1<<j) == 0 {
						next := mask | (1 << j)
						if dp[next][j] == -1 || dp[next][j] > dp[mask][i]+dist[i][j] {
							dp[next][j] = dp[mask][i] + dist[i][j]
						}
					}
				}
			}
		}
	}
	ret := -1
	finalMask := (1 << numBlock) - 1
	for i := 0; i < numBlock; i++ {
		if ret == -1 || ret > dp[finalMask][i]+dist[i][numBlock+1] {
			ret = dp[finalMask][i] + dist[i][numBlock+1]
		}
	}
	return ret
}

func bfs(x, y int, maze []string) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
		for j := 0; j < m; j++ {
			ret[i][j] = -1
		}
	}
	ret[x][y] = 0
	queue := [][]int{}
	queue = append(queue, []int{x, y})
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		curx, cury := p[0], p[1]
		for k := 0; k < 4; k++ {
			nx, ny := curx+dx[k], cury+dy[k]
			if inBound(nx, ny) && maze[nx][ny] != '#' && ret[nx][ny] == -1 {
				ret[nx][ny] = ret[curx][cury] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}
	return ret
}

func inBound(x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}
