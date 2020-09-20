package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}

	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	used := make([]bool, n)
	temp := []int{}
	var dfs func(int)
	dfs = func(start int) {
		ans = append(ans, append([]int(nil), temp...))
		for i := start; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			used[i] = true
			dfs(i + 1)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)
	return ans
}
