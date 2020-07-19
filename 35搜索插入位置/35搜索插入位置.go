package main

import (
	"fmt"
)

func main() {
	nums := []int{
		1, 3, 5, 8, 9, 12, 23, 44, 53, 56, 77,
	}
	target := 7
	fmt.Println(searchInsertOfficial(nums, target))
	//fmt.Println(sort.SearchInts(nums, target))
}

func searchInsert(nums []int, target int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			index = i
			break
		}
		if target < nums[i] {
			index = i
			break
		}
		if i == len(nums)-1 {
			index = i + 1
		}
	}
	return index
}

func searchInsertOfficial(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
