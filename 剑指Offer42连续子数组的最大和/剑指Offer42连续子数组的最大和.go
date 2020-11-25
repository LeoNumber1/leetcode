package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArrayYOUHUA(nums))
}

func maxSubArray0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ans int = nums[0]
	temp := ans
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			if temp < 0 {
				temp = nums[i]
			} else {
				temp += nums[i]
			}
		} else {
			if temp < nums[i] {
				temp = nums[i]
			} else {
				temp += nums[i]
			}
		}

		if temp > ans {
			ans = temp
		}
	}
	return ans
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp = make([]int, len(nums)) //定义dp[i],值为当前包含i在内的最大和，i为index
	var ans = nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		temp := dp[i]
		if dp[i] > ans {
			ans = temp
		}
	}
	return ans
}

func maxSubArrayYOUHUA(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var pre = nums[0]
	var ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if pre > 0 {
			pre += nums[i]
		} else {
			pre = nums[i]
		}
		if pre > ans {
			ans = pre
		}
	}
	return ans
}
