package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2, 3, 1}

	fmt.Println(reversePairsOfficial2(nums))
}

func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	n1 := append([]int(nil), nums[:n/2]...)
	n2 := append([]int(nil), nums[n/2:]...)
	cnt := reversePairs(n1) + reversePairs(n2) // 递归完毕后，n1 和 n2 均为有序

	// 统计重要翻转对 (i,j) 的数量
	// 由于 n1 和 n2 均为有序，可以用两个指针同时遍历
	j := 0
	for _, v := range n1 {
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		cnt += j
	}

	// n1 和 n2 归并填入 nums
	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt
}

type fenwick struct {
	tree []int
}

func newFenwickTree(n int) fenwick {
	return fenwick{make([]int, n+1)}
}

func (f fenwick) add(i int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i]++
	}
}

func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func reversePairsOfficial2(nums []int) (cnt int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 离散化所有下面统计时会出现的元素
	allNums := make([]int, 0, 2*n)
	for _, v := range nums {
		allNums = append(allNums, v, 2*v)
	}
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < 2*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	t := newFenwickTree(k)
	for i, v := range nums {
		// 统计之前插入了多少个比 2*v 大的数
		cnt += i - t.sum(kth[2*v])
		t.add(kth[v])
	}
	return
}
