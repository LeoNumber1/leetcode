package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 3}
	k := 2

	//nums = []int{1, 2}
	//k = 2

	//nums = []int{4, 1, -1, 2, -1, 2, 3}
	//k = 2

	nums = []int{5, 3, 1, 1, 1, 3, 73, 1}
	k = 2

	//fmt.Println(topKFrequent(nums, k))
	//fmt.Println(topKFrequentOfficial(nums, k))
	fmt.Println(topKFrequentOfficial1(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	arr := []int{}
	for _, v := range nums {
		m[v] += 1
		insert(&arr, v, m)
	}

	return arr[:k]
}

func insert(arr *[]int, val int, m map[int]int) {
	if len(*arr) == 0 {
		*arr = append(*arr, val)
		return
	}
	for i := 0; i < len(*arr); i++ {
		if val != (*arr)[i] {
			if m[val] < m[(*arr)[i]] {
				if i == len(*arr)-1 {
					*arr = append(*arr, val)
					return
				} else {
					continue
				}
			} else if m[val] == m[(*arr)[i]] {
				has := false
				for j := i + 1; j < len(*arr); j++ {
					if val == (*arr)[j] {
						has = true
						break
					}
				}
				if !has {
					*arr = append(*arr, val)
					return
				}
			} else {
				for j := i + 1; j < len(*arr); j++ {
					if val == (*arr)[j] {
						(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
						return
					}
				}
				*arr = append(*arr, val)
				return
			}
		} else {
			if i != 0 && m[val] > m[(*arr)[i-1]] {
				(*arr)[i], (*arr)[i-1] = (*arr)[i-1], (*arr)[i]
			}
			return
		}
	}
}

func topKFrequentOfficial(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentOfficial1(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	values := [][]int{}
	for key, value := range occurrences {
		values = append(values, []int{key, value})
	}
	ret := make([]int, k)
	qsort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		qsort(values, start, index-1, ret, retIndex, k)
	} else {
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}
