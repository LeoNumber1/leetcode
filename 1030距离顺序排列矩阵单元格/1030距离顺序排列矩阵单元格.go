package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(allCellsDistOrder(2, 3, 1, 2))
	//fmt.Println(allCellsDistOrder(89, 90, 21, 65))
	fmt.Println(allCellsDistOrderOfficial3(89, 90, 21, 65))
	//fmt.Println(allCellsDistOrderOfficial3(1, 2, 0, 0))
}

//执行耗时:20 ms,击败了91.80% 的Go用户
//内存消耗:8 MB,击败了25.00% 的Go用户
func allCellsDistOrder0(R int, C int, r0 int, c0 int) [][]int {
	var m = make(map[int][][]int)
	var min, max int = 201, 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dis := distance(i, j, r0, c0)
			if dis > max {
				max = dis
			}
			if dis < min {
				min = dis
			}
			m[dis] = append(m[dis], []int{i, j})
		}
	}

	var ans [][]int
	for i := min; i <= max; i++ {
		if v, has := m[i]; has {
			for _, val := range v {
				ans = append(ans, val)
			}
		}
	}
	return ans
}

//执行耗时:24 ms,击败了65.57% 的Go用户
//内存消耗:7.7 MB,击败了28.57% 的Go用户
func allCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	var arr = make([][][]int, 201)
	var min, max int = 201, 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dis := distance(i, j, r0, c0)
			if dis > max {
				max = dis
			}
			if dis < min {
				min = dis
			}
			arr[dis] = append(arr[dis], []int{i, j})
		}
	}

	var ans [][]int
	for i := min; i <= max; i++ {
		for _, v := range arr[i] {
			ans = append(ans, v)
		}
	}
	return ans
}

func distance(i, j, r0, c0 int) int {
	return abs(i-r0) + abs(j-c0)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var direction = [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

//执行耗时:24 ms,击败了65.57% 的Go用户
//内存消耗:7.2 MB,击败了53.57% 的Go用户
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	touch := make([][]bool, R)
	for k := range touch {
		touch[k] = make([]bool, C)
	}
	var ans [][]int
	queue := [][]int{{r0, c0}}
	touch[r0][c0] = true
	for len(queue) > 0 {
		this := queue[0]
		queue = queue[1:]
		ans = append(ans, this)
		for _, v := range direction {
			newR := this[0] + v[0]
			newC := this[1] + v[1]
			if newR < R && newR >= 0 && newC < C && newC >= 0 && !touch[newR][newC] {
				touch[newR][newC] = true
				queue = append(queue, []int{newR, newC})
			}
		}
	}
	return ans
}

//执行耗时:28 ms,击败了47.54% 的Go用户
//内存消耗:7 MB,击败了69.64% 的Go用户
func allCellsDistOrderOfficial1(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return ans
}

//执行耗时:20 ms,击败了91.80% 的Go用户
//内存消耗:7.1 MB,击败了66.07% 的Go用户
func allCellsDistOrderOfficial2(n, m, r0, c0 int) [][]int {
	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
	buckets := make([][][]int, maxDist+1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dist := abs(i-r0) + abs(j-c0)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}

	ans := make([][]int, 0, n*m)
	for _, bucket := range buckets {
		ans = append(ans, bucket...)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dir4 = [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}} //右上角、左上角，左下角、右下角

//执行耗时:16 ms,击败了96.72% 的Go用户
//内存消耗:7.1 MB,击败了66.07% 的Go用户
func allCellsDistOrderOfficial3(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 1, n*m)
	ans[0] = []int{r0, c0}
	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
	row, col := r0, c0
	for dist := 1; dist <= maxDist; dist++ {
		row--
		for i, dir := range dir4 {
			for i%2 == 0 && row != r0 || i%2 == 1 && col != c0 {
				if 0 <= row && row < n && 0 <= col && col < m {
					ans = append(ans, []int{row, col})
				}
				row += dir[0]
				col += dir[1]
			}
		}
	}
	return ans
}
