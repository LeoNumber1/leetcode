package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
	fmt.Println(intersectionOfficial1(nums1, nums2))
}

//4 ms	3 MB-57%
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	for _, v := range nums1 {
		m[v] = true
	}
	var ans []int
	for _, v := range nums2 {
		if _, has := m[v]; has {
			ans = append(ans, v)
			delete(m, v)
		}
	}
	return ans
}

//	0 ms-100.00%	3.2 MB-12.37%
func intersectionOfficial(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}

//	4 ms-88.35% 	2.8 MB-97.90%
func intersectionOfficial1(nums1 []int, nums2 []int) (ans []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if ans == nil || x > ans[len(ans)-1] {
				ans = append(ans, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}
