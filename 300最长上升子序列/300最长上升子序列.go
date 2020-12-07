package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums = []int{10, 9, 2, 5, 3, 7, 1, 18}

	fmt.Println(lengthOfLIS(nums))
}

//72 ms	3.8 MB
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n) //值为包含当前位置在内的最大上升序列的长度
	dp[0] = 1
	var ans int = 1
	for i := 1; i < n; i++ {
		temp := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > temp {
				temp = dp[j]
			}
		}
		dp[i] = temp + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
