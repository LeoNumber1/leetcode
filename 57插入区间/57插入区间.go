package main

import "fmt"

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	//单边在
	newInterval := []int{2, 5}
	newInterval = []int{4, 7}

	//双边都不在
	newInterval = []int{4, 5}

	//双边在同一个里
	newInterval = []int{7, 8}

	//双边在不同的里
	newInterval = []int{2, 7}

	//单边越界
	newInterval = []int{-1, 0}
	newInterval = []int{-1, 1}
	newInterval = []int{9, 10}
	newInterval = []int{10, 11}
	newInterval = []int{7, 10}
	newInterval = []int{5, 10}

	////双边越界
	//newInterval = []int{-1, 9}
	//newInterval = []int{-1, 10}
	//newInterval = []int{1, 10}

	//intervals = [][]int{{1, 5}}
	//newInterval = []int{2, 7}
	//newInterval = []int{0, 2}

	//跨多个
	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 18}

	//fmt.Println(insert(intervals, newInterval))
	fmt.Println(insertOfficial(intervals, newInterval))
}

//4 ms-99.40%	4.8 MB-40.43%
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}

	var ans [][]int
	var index = 0
	x, y := newInterval[0], newInterval[1]
	i := 0
	if y < intervals[0][0] {
		ans = append(ans, newInterval)
		ans = append(ans, intervals[:]...)
		return ans
	} else if y == intervals[0][0] {
		ans = append(ans, []int{newInterval[0], intervals[0][1]})
		ans = append(ans, intervals[1:]...)
		return ans
	}
	if x > intervals[n-1][1] {
		ans = append(ans, intervals[:]...)
		ans = append(ans, newInterval)
		return ans
	} else if x == intervals[n-1][1] {
		ans = append(ans, intervals[:n-1]...)
		ans = append(ans, []int{intervals[n-1][0], newInterval[1]})
		return ans
	}
	if x <= intervals[0][0] && y >= intervals[n-1][1] {
		ans = append(ans, newInterval)
		return ans
	}

	var temp []int
	for i < n {
		if x > intervals[i][1] {
			ans = append(ans, intervals[i])
			i++
			continue
		}
		if y < intervals[i][0] {
			if len(temp) == 0 {
				ans = append(ans, newInterval)
				index = i
				break
			} else {
				ans = append(ans, append(temp, y))
				index = i
				break
			}
		}
		if len(temp) == 0 {
			if x >= intervals[i][0] && x <= intervals[i][1] {
				temp = append(temp, intervals[i][0])
			} else {
				temp = append(temp, x)
			}
		}
		if y >= intervals[i][0] && y <= intervals[i][1] {
			if len(temp) != 0 {
				temp = append(temp, intervals[i][1])
				ans = append(ans, temp)
			} else {
				ans = append(ans, []int{x, intervals[i][1]})
			}
			index = i + 1
			break
		} else {
			if i == n-1 {
				temp = append(temp, y)
				ans = append(ans, temp)
				return ans
			}
			i++
			continue
		}
	}
	if index < n {
		ans = append(ans, intervals[index:]...)
	}
	return ans
}

func insertOfficial(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
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
