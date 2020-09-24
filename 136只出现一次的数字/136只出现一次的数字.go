package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	nums = []int{4, 1, 2, 1, 2}

	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var ans int
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
