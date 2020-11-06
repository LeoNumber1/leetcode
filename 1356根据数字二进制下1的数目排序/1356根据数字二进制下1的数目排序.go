package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//arr = []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	//arr = []int{2, 3, 5, 7, 11, 13, 17, 19}

	//fmt.Println(sortByBits(arr))
	fmt.Println(sortByBitsOfficial(arr))
	fmt.Println(sortByBitsOfficial1(arr))
}

//8 ms-94.05%	4.4 MB-35.44%
func sortByBits(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	m := map[int][]int{}
	var max1 int
	for _, v := range arr {
		num1 := count1(v)
		if num1 > max1 {
			max1 = num1
		}
		m[num1] = append(m[num1], v)
	}
	var ans []int
	for i := 0; i <= max1; i++ {
		if len(m[i]) > 0 {
			sort.Ints(m[i])
			ans = append(ans, m[i]...)
		}
	}
	return ans
}

func count1(a int) int {
	if a == 0 {
		return 0
	}
	var count int
	const BASE = 2
	for a != 0 {
		if a%BASE == 1 {
			count++
		}
		a /= BASE
	}

	return count
}

func onesCount(x int) (c int) {
	for ; x > 0; x /= 2 {
		c += x % 2
	}
	return
}

//12 ms-53.57%	3.5 MB-82.28%
func sortByBitsOfficial(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := onesCount(x), onesCount(y)
		return cx < cy || cx == cy && x < y
	})
	return a
}

var bit = [1e4 + 1]int{}

func init() {
	for i := 1; i <= 1e4; i++ {
		bit[i] = bit[i>>1] + i&1
	}
}

//12 ms-53.57%	3.6 MB-63.29%
func sortByBitsOfficial1(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	})
	return a
}
