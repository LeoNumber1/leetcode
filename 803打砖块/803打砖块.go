package main

import "fmt"

func main() {
	grid := [][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 1, 1, 0}}
	hits := [][]int{{2, 0}}

	fmt.Println(hitBricks(grid, hits))
}

var directions = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func hitBricks(grid [][]int, hits [][]int) []int {
	h, w := len(grid), len(grid[0])
	fa := make([]int, h*w+1)
	size := make([]int, h*w+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(from, to int) {
		from, to = find(from), find(to)
		if from != to {
			size[to] += size[from]
			fa[from] = to
		}
	}

	status := make([][]int, h)
	for i, row := range grid {
		status[i] = append([]int(nil), row...)
	}
	// 遍历 hits 得到最终状态
	for _, p := range hits {
		status[p[0]][p[1]] = 0
	}

	// 根据最终状态建立并查集
	root := h * w
	for i, row := range status {
		for j, v := range row {
			if v == 0 {
				continue
			}
			if i == 0 {
				union(i*w+j, root)
			}
			if i > 0 && status[i-1][j] == 1 {
				union(i*w+j, (i-1)*w+j)
			}
			if j > 0 && status[i][j-1] == 1 {
				union(i*w+j, i*w+j-1)
			}
		}
	}

	type pair struct{ x, y int }
	directions := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		p := hits[i]
		r, c := p[0], p[1]
		if grid[r][c] == 0 {
			continue
		}

		preSize := size[find(root)]
		if r == 0 {
			union(c, root)
		}
		for _, d := range directions {
			if newR, newC := r+d.x, c+d.y; 0 <= newR && newR < h && 0 <= newC && newC < w && status[newR][newC] == 1 {
				union(r*w+c, newR*w+newC)
			}
		}
		curSize := size[find(root)]
		if cnt := curSize - preSize - 1; cnt > 0 {
			ans[i] = cnt
		}
		status[r][c] = 1
	}
	return ans
}
