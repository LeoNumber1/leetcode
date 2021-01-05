package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	nums = []int{1, -9, 8, -6, 6, 4, 0, 5}
	k = 4

	t := time.Now()
	fmt.Println(maxSlidingWindow(nums, k), time.Since(t))
	//maxSlidingWindow(nums, k)
	//fmt.Println(time.Since(t))
}

type hp [][2]int

func (h hp) Len() int {
	return len(h)
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *hp) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func (h hp) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) push(x [2]int) {
	heap.Push(h, x)
}
func (h *hp) pop() [2]int {
	return heap.Pop(h).([2]int)
}
func (h hp) top() [2]int {
	return h[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 1 || k == 1 {
		return nums
	}
	left, right := 0, 0
	h := &hp{}
	heap.Init(h)
	for ; right < k; right++ {
		h.push([2]int{nums[right], right})
	}
	var ans = make([]int, n-k+1)
	var index = 1
	ans[0] = h.top()[0]
	for ; right < n; right++ {
		if nums[right] == nums[left] {
			ans[index] = h.top()[0]
			left++
			index++
			continue
		}
		for h.top()[1] <= left {
			h.pop()
		}
		h.push([2]int{nums[right], right})
		ans[index] = h.top()[0]
		left++
		index++
	}

	return ans
}

//超时了
func maxSlidingWindow0(nums []int, k int) []int {
	n := len(nums)
	if n == 1 || k == 1 {
		return nums
	}
	left, right := 0, 0
	slice := []int{}
	for ; right < k; right++ {
		slice = append(slice, nums[right])
	}
	var ans = make([]int, n-k+1)
	var index = 1
	sort.Ints(slice)
	ans[0] = slice[len(slice)-1]
	for ; right < n; right++ {
		if nums[right] == nums[left] {
			ans[index] = slice[len(slice)-1]
			left++
			index++
			continue
		}
		//开始二分删除
		l, r := 0, len(slice)-1
		key := 0
		newNum := nums[right]
		oldNum := nums[left]
		var temp []int
		if slice[l] == oldNum {
			slice = slice[1:]
			goto next1
		} else if slice[r] == oldNum {
			slice = slice[:len(slice)-1]
			goto next1
		}
		for l <= r {
			mid := l + (r-l)/2
			if slice[mid] == oldNum {
				key = mid
				break
			} else if slice[mid] > oldNum {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		slice = append(slice[:key], slice[key+1:]...)
	next1:
		//开始二分插入
		l, r = 0, len(slice)-1
		if newNum >= slice[r] {
			slice = append(slice, nums[right])
			goto next
		}
		if newNum <= slice[l] {
			slice = append([]int{newNum}, slice...)
			goto next
		}
		key = 0
		for l <= r {
			mid := l + (r-l)/2
			if newNum <= slice[mid] {
				key = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		temp = make([]int, len(slice)-key)
		copy(temp, slice[key:])
		slice[key] = newNum
		slice = append(slice[:key+1], temp...)
	next:
		ans[index] = slice[len(slice)-1]
		left++
		index++
	}

	return ans
}
