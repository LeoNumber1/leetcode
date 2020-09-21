package main

import "fmt"

func main() {
	target := []int{2, 4, 3, 1, 5}
	target = []int{5, 4, 3, 2, 1}
	target = []int{2, 4, 1, 5, 3}
	target = []int{2, 1, 4, 3, 5}
	target = []int{2, 1, 4, 5, 3}

	fmt.Println(isMagic(target))

	//target = []int{1, 2, 3, 4, 5}
	//fmt.Println(swap(target, 5))
}

func isMagic(target []int) bool {
	begin := []int{}
	n := len(target)
	for i := 0; i < n; i++ {
		begin = append(begin, i+1)
	}
	//开始交换
	begin = swap(begin, n)
	if begin[0] != target[0] {
		return false
	}
	var num int
	begin, target, num = compare(begin, target, n)
	for len(target) > 0 {
		begin = swap(begin, len(begin))
		if !compare1(begin, target, num) {
			return false
		}
		if len(target) >= num {
			begin = begin[num:]
			target = target[num:]
		} else {
			break
		}
	}
	return true
}

func swap(begin []int, n int) []int {
	a := []int{}
	b := []int{}
	for i := 0; i < n; i++ {
		if i%2 == 0 { //i是奇数
			a = append(a, begin[i])
		} else {
			b = append(b, begin[i])
		}
	}
	return append(b, a...)
}

func compare(begin, target []int, n int) ([]int, []int, int) {
	count := 0
	for i := 0; i < n; i++ {
		if begin[i] == target[i] {
			count++
		} else {
			break
		}
	}
	return begin[count:], target[count:], count
}

func compare1(begin, target []int, k int) bool {
	for i := 0; i < k && i < len(begin); i++ {
		if begin[i] != target[i] {
			return false
		}
	}
	return true
}
