package main

import "fmt"

func main() {
	nums := []int{
		0, 0, 1, 1, 1, 1, 2, 3, 3,
	}
	fmt.Println(removeDuplicates(nums))
	//fmt.Println(removeDuplicatesOfficial(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	var i, count int
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			count = 0
		} else {
			if count == 0 {
				i++
				nums[i] = nums[j]
				count++
			}
		}
	}
	return i + 1
}
