package main

import "fmt"

func main() {
	nums := []int{
		//1, 1, 2,
		0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
	}
	fmt.Println(removeDuplicates(nums))
	//fmt.Println(removeDuplicatesOfficial(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	var lenth int = len(nums)
	for k := 0; k < lenth; k++ {
		for k+1 < lenth {
			if nums[k] == nums[k+1] {
				if k+2 < lenth {
					nums = append(nums[:k+1], nums[k+2:]...)
				} else {
					nums = nums[:k+1]
				}
				lenth--
			} else {
				break
			}
		}
	}
	return len(nums)
}

func removeDuplicatesOfficial(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
