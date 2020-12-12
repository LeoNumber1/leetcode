package main

import "fmt"

func main() {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	//nums = []int{1, 2, 3, 1, 3, 4, 5, 5, 6, 4, 1}

	//fmt.Println(wiggleMaxLength(nums))
	fmt.Println(wiggleMaxLengthOfficial3(nums))
}

//0 ms-100.00%	2.2 MB-15.17%
func wiggleMaxLength0(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([][2]int, n) //
	dp[0][0] = 1            //包含当前数字在内的最大摆动序列长度
	dp[0][1] = 0            //此数与上个数之差，0为初始，1为正，-1为负
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			if dp[i-1][1] <= 0 {
				dp[i][0] = dp[i-1][0] + 1
			} else {
				dp[i][0] = dp[i-1][0]
			}
			dp[i][1] = 1
		} else if nums[i] < nums[i-1] {
			if dp[i-1][1] >= 0 {
				dp[i][0] = dp[i-1][0] + 1
			} else {
				dp[i][0] = dp[i-1][0]
			}
			dp[i][1] = -1
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1][0]
}

//0 ms-100.00%	2.1 MB-100.00%
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := [2]int{1, 0}
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			if dp[1] <= 0 {
				dp[0]++
			}
			dp[1] = 1
		} else if nums[i] < nums[i-1] {
			if dp[1] >= 0 {
				dp[0]++
			}
			dp[1] = -1
		}
	}
	return dp[0]
}

func wiggleMaxLengthOfficial1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(up[i-1]+1, down[i-1])
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func wiggleMaxLengthOfficial2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			//up = max(up, down+1)
			up = down + 1
		} else if nums[i] < nums[i-1] {
			//down = max(up+1, down)
			down = up + 1
		}
	}
	return max(up, down)
}

func wiggleMaxLengthOfficial3(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}
