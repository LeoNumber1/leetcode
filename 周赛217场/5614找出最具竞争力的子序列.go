package main

import "fmt"

func main() {
	nums := []int{2, 4, 3, 3, 5, 4, 9, 6}
	k := 4
	fmt.Println(mostCompetitive(nums, k))
}

func mostCompetitive0(nums []int, k int) []int {
	n := len(nums)
	if n <= k {
		return nums
	}
	ans := make([]int, 0)
	var minPos = -1
	for k > 0 {
		start := minPos + 1
		minVal := nums[start]
		minPos = start
		for i := start; i <= n-k; i++ {
			if nums[i] < minVal {
				minVal = nums[i]
				minPos = i
			}
			if minVal == 0 {
				break
			}
		}
		ans = append(ans, minVal)
		k--
	}
	return ans
}

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	if n <= k {
		return nums
	}
	ans := make([]int, 0)
	var minPos = -1
	for k > 0 {
		start := minPos + 1
		minVal := nums[start]
		minPos = start
		for i := start; i <= n-k; i++ {
			if nums[i] < minVal {
				minVal = nums[i]
				minPos = i
			}
			if minVal == 0 {
				break
			}
		}
		ans = append(ans, minVal)
		k--
	}
	return ans
}
