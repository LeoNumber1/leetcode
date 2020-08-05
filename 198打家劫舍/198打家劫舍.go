package main

import "fmt"

func main() {
	nums := []int{
		2, 7, 9, 3, 1,
		//1, 2, 3, 1,
		//3, 2, 3, 4, 2, 3,
	}

	fmt.Println(rob(nums))
	fmt.Println(robOfficial(nums))
}

func rob(nums []int) int {
	res := make([]int, len(nums))
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		res[i] = max(res[i-1], res[i-2]+nums[i])
	}
	return res[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robOfficial(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
