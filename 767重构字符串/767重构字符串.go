package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	s := "baaba"

	s = "bbaaa"
	//s = "aaabb"
	//s = "vvvlo"
	s = "baaac"
	s = "abbabbaaab"

	//fmt.Println(reorganizeString(s))
	fmt.Println(reorganizeString1(s))
}

//0 ms-100.00%	2.2 MB-18.18%
func reorganizeString0(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var arr = make([]int, 26)
	for _, v := range s {
		arr[v-'a']++
	}
	var res []byte
	for n > 0 {
		conti := false
		for k, v := range arr {
			if v != 0 {
				val := byte('a' + k)
				if len(res) == 0 || val != res[len(res)-1] {
					res = append(res, val)
					arr[k]--
					n--
					conti = true
				} else {
					if byte('a'+k) != res[0] {
						res = append([]byte{val}, res...)
						arr[k]--
						n--
						conti = true
						continue
					}
					for i := 1; i < len(res); i++ {
						if i+1 < len(res) && res[i] != val && res[i+1] != val {
							temp := make([]byte, len(res[i+1:]))
							copy(temp, res[i+1:])
							res = append(append(res[:i+1], val), temp...)
							arr[k]--
							n--
							conti = true
							break
						}
					}
					if !conti {
						return ""
					}
				}
			}
		}
		if !conti {
			return ""
		}
	}
	return string(res)
}

//双指针 0 ms-100.00%	2 MB-95.45%
func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var arr = []byte(s)
	ans := reorg(arr, n)
	if len(ans) > 0 {
		return ans
	}
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return reorg(arr, n)
}

func reorg(arr []byte, n int) string {
	var slow, fast = 0, 1
	for fast < n {
		if arr[slow] != arr[fast] {
			slow++
			fast++
			continue
		}
		slow = fast
		for fast < n && arr[slow] == arr[fast] {
			fast++
		}
		if fast == n {
			return ""
		}
		arr[slow], arr[fast] = arr[fast], arr[slow]
		slow++
		fast = slow + 1
	}
	return string(arr)
}

type myHeap [][]int

func (hp *myHeap) Push(v interface{}) {
	*hp = append(*hp, v.([]int))
}

func (hp *myHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	x := old[n-1]
	*hp = old[:n-1]
	return x
}

func (hp myHeap) Less(i, j int) bool {
	return hp[i][1] > hp[j][1] || hp[i][1] == hp[j][1] && hp[i][0] < hp[j][0]
}
func (hp myHeap) Swap(i, j int) { hp[i], hp[j] = hp[j], hp[i] }
func (hp myHeap) Len() int      { return len(hp) }

//0 ms-100.00%	2.3 MB-10.66%
func reorganizeString1(s string) string {
	n := len(s)
	m := make(map[int32]int)
	for _, ch := range s {
		m[ch-'a']++
	}
	maxCh := 0
	hp := &myHeap{}
	for k, v := range m {
		if v > maxCh {
			maxCh = v
		}
		hp.Push([]int{int(k), v})
	}
	if maxCh > (n+1)>>1 {
		return ""
	}
	heap.Init(hp)

	var ans []byte
	for hp.Len() > 0 {
		var most, more []int
		most = heap.Pop(hp).([]int)
		ans = append(ans, byte(most[0]+'a'))
		if hp.Len() > 0 {
			more = heap.Pop(hp).([]int)
			ans = append(ans, byte(more[0]+'a'))
			more[1]--
			if more[1] > 0 {
				heap.Push(hp, more)
			}
		}
		most[1]--
		if most[1] > 0 {
			heap.Push(hp, most)
		}
	}

	return string(ans)
}

var cnt [26]int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return cnt[h.IntSlice[i]] > cnt[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func reorganizeStringOfficial(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	cnt = [26]int{}
	maxCnt := 0
	for _, ch := range s {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	h := &hp{}
	for i, c := range cnt[:] {
		if c > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)

	ans := make([]byte, 0, n)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		ans = append(ans, byte('a'+i), byte('a'+j))
		if cnt[i]--; cnt[i] > 0 {
			h.push(i)
		}
		if cnt[j]--; cnt[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}
	return string(ans)
}
