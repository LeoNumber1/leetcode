package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}

	fmt.Println(singleNumbers(nums))
}

func singleNumbers(nums []int) []int {
	var num int
	for _, v := range nums {
		num ^= v
	}

	a := num & (-num) //找到最右边的1
	var ans1, ans2 int

	for _, v := range nums {
		if a&v == 0 {
			ans1 ^= v
		} else {
			ans2 ^= v
		}
	}

	return []int{ans1, ans2}
}
