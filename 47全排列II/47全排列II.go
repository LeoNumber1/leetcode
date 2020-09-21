package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{1, 1, 2, 2}

	//fmt.Println(permuteUnique(nums))
	fmt.Println(permuteUniqueOfficial(nums))
}

//188 ms-5.12%	8.6 MB-5.03%
func permuteUnique(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n == 0 {
		return ans
	} else if n == 1 {
		return append(ans, nums)
	}
	used := map[int]bool{}
	has := map[string]bool{}

	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		if index == n {
			var s string
			for i := 0; i < n; i++ {
				s += strconv.Itoa(tmp[i]) + ","
			}
			if !has[s] {
				ans = append(ans, append([]int(nil), tmp...))
				has[s] = true
			}
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				tmp = append(tmp, nums[i])
				used[i] = true
				dfs(index+1, tmp)
				delete(used, i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0, []int{})
	return ans
}

func permuteUniqueOfficial(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
