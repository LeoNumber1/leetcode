package main

import "fmt"

func main() {
	nums := []int{
		//0, 2, 3, 4, 5,
		//1, 1, 1,
		-2, -2, -1, 3, 4, 5,
	}
	fmt.Println(findMagicIndex(nums))
	fmt.Println(findMagicIndexOfficial(nums))
}

func findMagicIndex(nums []int) int {
	for k, v := range nums {
		if k == v {
			return k
		}
	}
	return -1
}

func findMagicIndexOfficial(nums []int) int {
	return getAnswer(nums, 0, len(nums)-1)
}

func getAnswer(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (right-left)/2 + left
	leftAnswer := getAnswer(nums, left, mid-1)
	if leftAnswer != -1 {
		return leftAnswer
	} else if nums[mid] == mid {
		return mid
	}
	return getAnswer(nums, mid+1, right)
}
