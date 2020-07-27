package main

import (
	"fmt"
)

func main() {
	nums := []int{
		7, 2, 5, 10, 8,
	}
	m := 2
	fmt.Println(splitArrayOfficial(nums, m))
}

func splitArray(nums []int, m int) int {
	var max1, max2, sum, res int
	//var spArr1 []int
	var maxArr []int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if m == 1 {
		return sum
	}

	for i := 0; i < len(nums); i++ {
		max1 += nums[i]
		max2 = sum - max1
		maxArr = append(maxArr, max(max1, max2))
	}
	for k, v := range maxArr {
		if k == 0 {
			res = v
		} else {
			res = min(res, v)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func splitArrayOfficial(nums []int, m int) int {
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		right += nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}
	for left < right {
		mid := (right-left)/2 + left
		if check(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(nums []int, x, m int) bool {
	sum, cnt := 0, 1
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > x {
			cnt++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return cnt <= m
}
