package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	type args struct {
		n           int
		connections string
	}

	tests := []struct {
		index  int
		args   args
		target string
	}{
		{1, args{3, "[[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]]"}, "[2,0,3]"},
		{2, args{6, "[[1,1,0],[1,1,0],[1,1,1],[1,1,1],[0,0,0],[1,1,1],[1,0,0]]"}, "[4,6,0,1,2,3]"},
	}

	str2matrix := func(s string) (matrix [][]int) {
		arr := strings.Split(s, "],")
		for _, s2 := range arr {
			s2 = strings.TrimLeft(s2, "[")
			s2 = strings.TrimRight(s2, "]")
			arr1 := strings.Split(s2, ",")
			var temp []int
			for _, s3 := range arr1 {
				i, _ := strconv.Atoi(s3)
				temp = append(temp, i)
			}
			matrix = append(matrix, temp)
		}
		return
	}

	arr2str := func(arr []int) string {
		n := len(arr)
		if n == 0 {
			return "[]"
		}
		s := "["
		for k, v := range arr {
			s += strconv.Itoa(v)
			if k < n-1 {
				s += ","
			} else {
				s += "]"
			}
		}
		return s
	}

	var errNum bool
	for _, tt := range tests {
		result := kWeakestRowsOfficial(str2matrix(tt.args.connections), tt.args.n)
		if tt.target != arr2str(result) {
			errNum = true
			fmt.Println("—————— err in index:", tt.index, "except:", tt.target, " get result:", result)
		}
	}

	if !errNum {
		fmt.Println("------- All tests are OK! -------")
	}
}

//最简单暴力法+排序，时间 O(m * n), 空间 O(m)    12 ms	4.9 MB
func kWeakestRows0(mat [][]int, k int) []int {
	n := len(mat)
	stack := make([][2]int, n)
	for k, v := range mat {
		count := 0
		for _, value := range v {
			if value == 1 {
				count++
			} else {
				break
			}
		}
		stack[k] = [2]int{k, count}
	}
	sort.Slice(stack, func(i, j int) bool {
		return stack[i][1] < stack[j][1] || (stack[i][1] == stack[j][1] && stack[i][0] < stack[j][0])
	})
	var ans = make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = stack[i][0]
	}
	return ans
}

type myHeap [][2]int

func (h myHeap) Len() int {
	return len(h)
}
func (h myHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1] || (h[i][1] == h[j][1] && h[i][0] < h[j][0])
}
func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *myHeap) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// 小顶堆  8 ms	4.9 MB
func kWeakestRows(mat [][]int, k int) []int {
	h := &myHeap{}
	heap.Init(h)
	for k, v := range mat {
		count := 0
		for _, value := range v {
			if value == 1 {
				count++
			} else {
				break
			}
		}
		heap.Push(h, [2]int{k, count})
	}
	var ans = make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

//官方堆 + 二分查找    8 ms	4.9 MB
func kWeakestRowsOfficial(mat [][]int, k int) []int {
	h := hp{}
	for i, row := range mat {
		pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		h = append(h, pair{pow, i})
	}
	heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).idx
	}
	return ans
}

type pair struct{ pow, idx int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
