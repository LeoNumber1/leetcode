package main

import "fmt"

func main() {
	nums := []int{1, 3}
	n := 6
	fmt.Println(minPatches(nums, n))
}

func minPatches(nums []int, n int) (patches int) {
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			patches++
		}
	}
	return
}
