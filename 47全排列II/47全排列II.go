package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{1, 3, 3, 2, 2}

	fmt.Println(len(permuteUnique(nums)))
	fmt.Println(len(permuteUnique1(nums)))
	fmt.Println(len(permuteUniqueOfficial(nums)))
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

//这个是插入法的进阶，但是太复杂了，搞了一半没搞出来，以后有时间再说
func permuteUnique1(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n == 0 {
		return ans
	} else if n == 1 {
		return append(ans, nums)
	}
	sort.Ints(nums)
	ans = [][]int{
		{nums[0]},
	}

	for k := 1; k < n; k++ {
		num := nums[k]
		ansLen := len(ans)
		for i := 0; i < ansLen; i++ {
			var index int
			if nums[k] == nums[k-1] {
				index = i
			}
			for j := index; j <= len(ans[i]); j++ {
				step := 0 //跨越步长，即后面有几个同样的数字
				tmp := []int{}
				if j == 0 {
					tmp = append(append(tmp, num), ans[i]...)
				} else {
					t := make([]int, len(ans[i][:j]))
					copy(t, ans[i][:j])
					tmp = append(append(t, num), ans[i][j:]...)
				}
				for z := j; z < len(ans[i]); z++ {
					if ans[i][z] == num {
						step++
					} else {
						break
					}
				}
				j += step
				ans = append(ans, tmp)
			}
		}
		ans = ans[ansLen:]
	}

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
