package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 2, 1}
	k := 2

	arr = []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	k = 8

	fmt.Println(getLeastNumbers(arr, k))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

//大顶堆实现 52 ms	6.6 MB
func getLeastNumbers(arr []int, k int) []int {
	h := &IntHeap{}
	for i := 0; i < k; i++ {
		*h = append(*h, arr[i])
	}

	heap.Init(h)

	for i := k; i < len(arr); i++ {
		heap.Push(h, arr[i])
		heap.Pop(h)
	}

	var ans = make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(h).(int)
	}

	return ans
}

//40 ms	6.5 MB
func getLeastNumbers0(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
