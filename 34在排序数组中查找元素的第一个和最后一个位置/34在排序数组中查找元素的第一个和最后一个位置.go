package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	nums = []int{1}
	target = 0

	fmt.Println(searchRangeOfficial(nums, target))
}

//我的二分法
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	left, right := 0, n-1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			break
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}
	ans := [2]int{}
	i, j := index, index
	for ; j >= 0; j-- {
		if nums[j] != target {
			break
		}
		ans[0] = j
	}
	for ; i < n; i++ {
		if nums[i] != target {
			break
		}
		ans[1] = i
	}
	return ans[:]
}

//官方二分法
func searchRangeOfficial(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
