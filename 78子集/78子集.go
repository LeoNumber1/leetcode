package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	//nums = []int{1, 2, 3, 4}

	//fmt.Println(subsets(nums))
	//fmt.Println(subsetsOfficial(nums))
	fmt.Println(subsetsOfficial2(nums))
}

//0 ms	2.2 MB
func subsets(nums []int) [][]int {
	ans := [][]int{{}}
	n := len(nums)
	temp := []int{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			return
		}
		for i := index; i < n; i++ {
			temp = append(temp, nums[i])
			ans = append(ans, append([]int(nil), temp...))
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)
	return ans
}

//0 ms	2.3 MB
func subsetsOfficial(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return ans
}

//0 ms	2.2 MB
func subsetsOfficial2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
