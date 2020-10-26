package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8, 1, 2, 2, 3}

	//fmt.Println(smallerNumbersThanCurrent(nums))
	//fmt.Println(smallerNumbersThanCurrentOfficial2(nums))
	fmt.Println(smallerNumbersThanCurrentOfficial3(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	data := make([][2]int, n)
	for key, num := range nums {
		data[key][0] = num
		data[key][1] = key
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	ans := make([]int, n)
	prev := -1
	for i, d := range data {
		if prev == -1 || d[0] != data[i-1][0] {
			prev = i
		}
		ans[d[1]] = prev
	}
	return ans
}

type pair struct {
	v   int
	pos int
}

func smallerNumbersThanCurrentOfficial2(nums []int) []int {
	n := len(nums)
	data := make([]pair, n)
	for key, num := range nums {
		data[key] = pair{num, key}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].v < data[j].v
	})
	ans := make([]int, n)
	prev := -1
	for i, d := range data {
		if prev == -1 || d.v != data[i-1].v {
			prev = i
		}
		ans[d.pos] = prev
	}
	return ans
}

func smallerNumbersThanCurrentOfficial3(nums []int) []int {
	cnt := [101]int{}
	maxNum := 0
	for _, num := range nums {
		cnt[num]++
		if num > maxNum {
			maxNum = num
		}
	}
	for i := 0; i < maxNum; i++ {
		cnt[i+1] += cnt[i]
	}
	ans := make([]int, len(nums))
	for k, num := range nums {
		if num > 0 {
			ans[k] = cnt[num-1]
		}
	}
	return ans
}
