package main

import (
	"fmt"
	"strconv"
)

func main() {
	grid := [][]int{
		{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0},
		//{1, 1}, {1, 1},
	}

	fmt.Println(islandPerimeter(grid))
	fmt.Println(grid)
}

var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

//76 ms	7.7 MB
func islandPerimeter1(grid [][]int) int {
	var ans int
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = 9
		for _, v := range direction {
			newI, newJ := i+v[0], j+v[1]
			if newI >= 0 && newI < n && newJ >= 0 && newJ < m { //值是有效的
				if grid[newI][newJ] == 0 {
					ans++
				} else if grid[newI][newJ] == 1 {
					dfs(newI, newJ)
				}
			} else {
				ans++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				goto end
			}
		}
	}
end:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 9 {
				grid[i][j] = 1
			}
		}
	}
	return ans
}

//136 ms	7.5 MB
func islandPerimeter2(grid [][]int) int {
	var ans int
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	indexMap := make(map[string]bool)
	arr := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				key := strconv.Itoa(i) + "," + strconv.Itoa(j)
				indexMap[key] = true
				arr = append(arr, []int{i, j})
				goto end
			}
		}
	}
end:
	for len(arr) > 0 {
		this := arr[0]
		arr = arr[1:]
		for _, v := range direction {
			newI, newJ := this[0]+v[0], this[1]+v[1]
			if newI >= 0 && newI < n && newJ >= 0 && newJ < m { //值是有效的
				if grid[newI][newJ] == 0 {
					ans++
				} else if grid[newI][newJ] == 1 {
					key := strconv.Itoa(newI) + "," + strconv.Itoa(newJ)
					if _, has := indexMap[key]; !has {
						indexMap[key] = true
						arr = append(arr, []int{newI, newJ})
					}
				}
			} else {
				ans++
			}
		}
	}

	return ans
}

//72 ms-83.33%	7 MB-19.49%
func islandPerimeter3(grid [][]int) int {
	var ans int
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	used := make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				used[i][j] = true
				for _, v := range direction {
					newI, newJ := i+v[0], j+v[1]
					if newI >= 0 && newI < n && newJ >= 0 && newJ < m { //值是有效的
						if grid[newI][newJ] == 0 {
							ans++
						}
					} else {
						ans++
					}
				}
			}
		}
	}

	return ans
}

//92 ms	6.9 MB
func islandPerimeter(grid [][]int) int {
	var ans int
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				for _, v := range direction {
					newI, newJ := i+v[0], j+v[1]
					if newI < 0 || newI >= n || newJ < 0 || newJ >= m || grid[newI][newJ] == 0 {
						ans++
					}
				}
			}
		}
	}

	return ans
}
