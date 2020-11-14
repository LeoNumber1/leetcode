package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}

	//fmt.Println(relativeSortArray(arr1, arr2))
	fmt.Println(relativeSortArrayOfficial2(arr1, arr2))
}

//0 ms-100.00%	2.4 MB-31.71%
func relativeSortArray(arr1 []int, arr2 []int) []int {
	if len(arr2) == 0 {
		sort.Ints(arr1)
		return arr1
	}
	var ans []int
	var arr = make([]int, len(arr2))
	var m = make(map[int]int)
	for k, v := range arr2 {
		m[v] = k
	}

	var temp []int
	for _, v := range arr1 {
		if val, has := m[v]; has {
			arr[val]++
		} else {
			temp = append(temp, v)
		}
	}
	for k, v := range arr {
		for i := 0; i < v; i++ {
			ans = append(ans, arr2[k])
		}
	}
	if len(temp) > 0 {
		sort.Ints(temp)
		ans = append(ans, temp...)
	}

	return ans
}

//0 ms-100.00%	2.4 MB-63.41%
func relativeSortArrayOfficial(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i, v := range arr2 {
		rank[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}

//0 ms-100.00%	2.4 MB-63.41%
func relativeSortArrayOfficial2(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)
	for _, v := range arr1 {
		frequency[v]++
	}

	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}
