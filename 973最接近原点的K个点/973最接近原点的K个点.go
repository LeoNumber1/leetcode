package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	points := [][]int{{3, 3}, {-2, 4}, {3, 2}, {5, -1}}
	K := 2

	//fmt.Println(kClosest(points, K))
	//fmt.Println(kClosestOfficial(points, K))
	//fmt.Println(kClosestOfficial1(points, K))
	fmt.Println(kClosestOfficial2(points, K))
}

//152 ms-21.52%	8 MB-39.74%
func kClosest0(points [][]int, K int) [][]int {
	n := len(points)
	if K >= n {
		return points
	}
	var ans [][]int
	var arr = make([][2]int, 0)
	//insertArr := func(index, val int) {
	//	length := len(arr)
	//	if length == 0 {
	//		arr = append(arr, [2]int{index, val})
	//		return
	//	}
	//	left, right := 0, length-1
	//	for left < right {
	//		mid := (left + right) / 2
	//		if arr[mid][1] > val {
	//			right = mid
	//		} else {
	//			left = mid
	//		}
	//	}
	//	temp := arr[left+1:]
	//	arr[left] = [2]int{index, val}
	//	arr = append(arr[:left+1], temp...)
	//}
	for k, v := range points {
		temp := v[0]*v[0] + v[1]*v[1]
		arr = append(arr, [2]int{k, temp})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1] || arr[i][1] < arr[j][1] && i < j
	})
	for i := 0; i < K; i++ {
		ans = append(ans, points[arr[i][0]])
	}

	return ans
}

//二分法插入数组   1096 ms-6.33%	11.5 MB-5.13%
func kClosest(points [][]int, K int) [][]int {
	n := len(points)
	if K >= n {
		return points
	}
	var ans [][]int
	var arr = make([][2]int, 0)
	insertArr := func(index, val int) {
		length := len(arr)
		if length == 0 {
			arr = append(arr, [2]int{index, val})
			return
		}
		left, right := 0, length-1
		var position int
		for left <= right {
			mid := (left + right) / 2
			if arr[mid][1] > val {
				right = mid - 1
			} else {
				left = mid + 1
			}
			position = left
		}
		if position >= length {
			arr = append(arr, [2]int{index, val})
			return
		}
		le := len(arr[position:])
		temp := make([][2]int, le)
		copy(temp, arr[position:])
		arr[position] = [2]int{index, val}
		arr = append(arr[:position+1], temp...)
	}
	for k, v := range points {
		temp := v[0]*v[0] + v[1]*v[1]
		insertArr(k, temp)
	}
	for i := 0; i < K; i++ {
		ans = append(ans, points[arr[i][0]])
	}

	return ans
}

//140 ms-39.24%	8.1 MB-38.46%
func kClosestOfficial(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}

type pair struct {
	dist  int
	point []int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kClosestOfficial1(points [][]int, k int) (ans [][]int) {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h) // O(k) 初始化堆
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	for _, p := range h {
		ans = append(ans, p.point)
	}
	return
}

func less(p, q []int) bool {
	return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
}

func kClosestOfficial2(points [][]int, k int) (ans [][]int) {
	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})
	fmt.Println(points)

	var quickSelect func(left, right int)
	quickSelect = func(left, right int) {
		if left == right {
			return
		}
		pivot := points[right] // 取当前区间 [left,right] 最右侧元素作为切分参照
		lessCount := left
		for i := left; i < right; i++ {
			if less(points[i], pivot) {
				points[i], points[lessCount] = points[lessCount], points[i]
				lessCount++
			}
		}
		// 循环结束后，有 lessCount 个元素比 pivot 小
		// 把 pivot 交换到 points[lessCount] 的位置
		// 交换之后，points[lessCount] 左侧的元素均小于 pivot，points[lessCount] 右侧的元素均不小于 pivot
		points[right], points[lessCount] = points[lessCount], points[right]
		if lessCount+1 == k {
			return
		} else if lessCount+1 < k {
			quickSelect(lessCount+1, right)
		} else {
			quickSelect(left, lessCount-1)
		}
	}
	quickSelect(0, len(points)-1)
	return points[:k]
}
