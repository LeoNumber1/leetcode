package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		tmp := target - v
		if val, ok := m[tmp]; ok {
			return []int{val, k}
		}
		m[v] = k
	}
	return []int{}
}
