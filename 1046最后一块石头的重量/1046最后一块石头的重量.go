package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	stones := []int{2, 7, 4, 1, 8, 1} //1
	//stones = []int{3, 7, 2}           //2
	//stones = []int{10, 4, 2, 10} //2
	//stones = []int{5, 1, 8, 10, 7}                                                                                                                                     //1
	stones = []int{434, 667, 378, 919, 212, 902, 240, 257, 208, 996, 411, 222, 557, 634, 425, 949, 755, 833, 785, 886, 40, 159, 932, 157, 764, 916, 85, 300, 130, 278} //1

	fmt.Println(lastStoneWeight(stones))
}

//0 ms	2.1 MB
func lastStoneWeight(stones []int) int {
	n := len(stones)
	sort.Ints(stones)
	for n > 1 {
		max, second := stones[n-1], stones[n-2]
		stones = stones[:n-2]
		n -= 2
		if max == second {
			continue
		}
		newStone := max - second
		if n == 0 {
			return newStone
		}
		if n == 1 {
			if stones[0] == newStone {
				return 0
			}
			if stones[0] > newStone {
				return stones[0] - newStone
			}
			return newStone - stones[0]
		}
		left, right := 0, n-1
		n += 1
		if newStone >= stones[right] {
			stones = append(stones, newStone)
			continue
		}
		if newStone <= stones[0] {
			stones = append([]int{newStone}, stones...)
			continue
		}
		index := 0
		for left <= right {
			mid := left + (right-left)/2
			if newStone <= stones[mid] {
				index = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		temp := make([]int, n-index)
		copy(temp, stones[index:])
		stones[index] = newStone
		stones = append(stones[:index+1], temp...)
	}
	if n == 1 {
		return stones[0]
	}
	return 0
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

//0 ms	2.0 MB
func lastStoneWeightOfficial(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
