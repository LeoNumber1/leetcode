package main

import "fmt"

func main() {
	J := "aA"
	S := "aAAbbbb"

	fmt.Println(numJewelsInStones1(J, S))
}

//0 ms-100.00%	2.1 MB-15.38%
func numJewelsInStones(J string, S string) int {
	m := make(map[int32]bool)
	for _, v := range J {
		m[v] = true
	}
	var ans int
	for _, v := range S {
		if m[v] {
			ans++
		}
	}
	return ans
}

//bitmap位运算	0 ms-100.00%	2 MB-55.62%
func numJewelsInStones1(J string, S string) int {
	bit := [4]int{}
	var ans int
	for _, v := range J {
		bit[v/32] |= 1 << (v % 32)
	}

	for _, v := range S {
		if bit[v/32]&(1<<(v%32)) != 0 {
			ans++
		}
	}
	return ans
}
