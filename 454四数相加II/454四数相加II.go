package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}

	fmt.Println(fourSumCount(A, B, C, D))
}

func fourSumCount0(A []int, B []int, C []int, D []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	var m1, m2 = make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := A[i], B[j]
			w, z := C[i], D[j]
			m1[x+y]++
			m2[w+z]++
		}
	}
	var arr []int
	for k, _ := range m2 {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	var ans int
	//开始二分
	for k, v := range m1 {
		left, right := 0, len(arr)-1
		for left <= right {
			mid := left + (right-left)/2
			temp := arr[mid] + k
			if temp == 0 {
				ans += v * m2[arr[mid]]
				break
			} else if temp < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return ans
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	var m1, m2 = make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := A[i], B[j]
			w, z := C[i], D[j]
			m1[x+y]++
			m2[w+z]++
		}
	}
	var ans int
	//开始二分
	for k, v := range m1 {
		if val, has := m2[-k]; has {
			ans += v * val
		}
	}

	return ans
}
