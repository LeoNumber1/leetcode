package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	nums = []int{3, 1, 3, 4, 3, 2, 3, 2}
	k = 6

	fmt.Println(maxOperations(nums, k))
}

func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	var ans int
	for i, v := range m {
		if val, has := m[k-i]; has {
			if i != k-i {
				ans += min(v, val)
			} else {
				ans += val / 2
			}
			delete(m, i)
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
