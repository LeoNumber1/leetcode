package main

import "fmt"

func main() {
	nums := []int{4, 2, 3}
	nums = []int{3, 4, 2, 3}
	nums = []int{5, 7, 1, 8}

	fmt.Println(checkPossibilityOfficial(nums))
}

func checkPossibilityOfficial(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
