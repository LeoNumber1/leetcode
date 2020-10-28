package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}

	fmt.Println(uniqueOccurrences(arr))
}

//0 ms	2.2 MB———— 双 100%
func uniqueOccurrences(arr []int) bool {
	m1 := make([]int, 2002)
	for _, v := range arr {
		m1[v+1000]++
	}
	m2 := make([]bool, 1001)
	for _, v := range m1 {
		if m2[v] {
			return false
		}
		if v != 0 {
			m2[v] = true
		}
	}
	return true
}

//0 ms	2.3 MB
func uniqueOccurrences2(arr []int) bool {
	m1 := map[int]int{}
	for _, v := range arr {
		m1[v]++
	}
	m2 := make([]bool, 1001)
	for _, v := range m1 {
		if m2[v] {
			return false
		}
		m2[v] = true
	}
	return true
}

func uniqueOccurrences1(arr []int) bool {
	m1 := map[int]int{}
	for _, v := range arr {
		m1[v]++
	}
	m2 := map[int]bool{}
	for _, v := range m1 {
		if m2[v] {
			return false
		}
		m2[v] = true
	}
	return true
}
