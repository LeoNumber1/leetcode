package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 12, -5, -6, 50, 105}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	var ans = float64(math.MinInt64)
	var temp int
	for i := 0; i < k; i++ {
		temp += nums[i]
	}
	floatTemp := float64(temp) / float64(k)
	if floatTemp > ans {
		ans = floatTemp
	}
	for i := k; i < len(nums); i++ {
		temp += nums[i] - nums[i-k]
		floatTemp = float64(temp) / float64(k)
		if floatTemp > ans {
			ans = floatTemp
		}
	}
	return ans
}
