package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}

//28 ms	6 MB
func majorityElement0(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//16 ms-98.27%	6.1 MB-41.09%
func majorityElement(nums []int) int {
	val := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == val {
			count++
		} else {
			if count == 1 {
				val = nums[i]
			} else {
				count--
			}
		}
	}
	return val
}
