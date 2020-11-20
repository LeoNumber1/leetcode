package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 4, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	nums2 = []int{7, 8, 9}
	//nums2 = []int{0, 1, 6}
	nums2 = []int{-3, -2, 3}
	n := 3
	//merge(nums1, m, nums2, n)
	mergeOfficial3(nums1, m, nums2, n)

	fmt.Println(nums1)
}

//0 ms-100.00%	2.3 MB-52.04%
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i := 0
	for i < m {
		if nums1[i] <= nums2[0] {
			i++
			continue
		}
		nums1[i], nums2[0] = nums2[0], nums1[i]
		for w := 0; w < n-1 && nums2[w] > nums2[w+1]; w++ {
			nums2[w], nums2[w+1] = nums2[w+1], nums2[w]
		}
		i++
	}
	length := len(nums1)
	j := 0
	for i := m; i < length; i++ {
		nums1[i] = nums2[j]
		j++
	}
}

//双指针，从后往前
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.3 MB,击败了52.04% 的Go用户
func mergeOfficial3(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1) - 1
	i, j := m-1, n-1
	for j >= 0 && i >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[index] = nums2[j]
			j--
		} else {
			nums1[index] = nums1[i]
			i--
		}
		index--
	}
	if j >= 0 {
		for q := index; q >= 0; q-- {
			nums1[q] = nums2[j]
			j--
		}
	}
}
