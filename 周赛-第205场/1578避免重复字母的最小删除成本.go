package main

import (
	"fmt"
	"time"
)

func main() {
	s := "abaac"
	cost := []int{1, 2, 3, 4, 5}

	//s = "abc"
	//cost = []int{1, 2, 3}

	//s = "aabaaaa"
	//cost = []int{1, 2, 3, 4, 1, 5, 2}

	s = "aaabbbabbbb"
	cost = []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}

	//t0 := time.Now()
	//fmt.Println(minCost(s, cost), time.Since(t0))
	//t1 := time.Now()
	//fmt.Println(minCost1(s, cost), time.Since(t1))
	//t2 := time.Now()
	//fmt.Println(minCost2(s, cost), time.Since(t2))
	t3 := time.Now()
	fmt.Println(minCost3(s, cost), time.Since(t3))
}

//312 ms	9 MB
func minCost(s string, cost []int) int {
	if len(s) < 2 {
		return 0
	}
	var ans int

	i := 0
	j := i + 1
	for i < len(s) && j < len(s) {
		if s[i] == s[j] {
			if cost[i] > cost[j] {
				ans += cost[j]
			} else {
				ans += cost[i]
				i = j
			}
			j++
		} else {
			i = j
			j++
		}
	}

	return ans
}

//232 ms	8.9 MB
func minCost1(s string, cost []int) int {
	if len(s) < 2 {
		return 0
	}
	var ans, tmp, max int
	var pre byte

	for i := 0; i <= len(s); i++ {
		if i == len(s) {
			ans += tmp - max
			break
		}
		if s[i] != pre {
			ans += tmp - max
			tmp = 0
			max = 0
			if i+1 < len(s) && s[i] == s[i+1] {
				tmp += cost[i]
				max = cost[i]
			}
			if i < len(s) {
				pre = s[i]
			}
		} else {
			tmp += cost[i]
			max = myMax(max, cost[i])
		}
	}

	return ans
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//160 ms	9 MB
func minCost2(s string, cost []int) int {
	st, et, rs, tmpMax, tmpSum := -1, -1, 0, 0, 0
	s += "*"
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if st == -1 {
				st = i
			}
			et = i + 1
			tmpSum += cost[i]
			if cost[i] > tmpMax {
				tmpMax = cost[i]
			}
			continue
		}
		if st != -1 && et != -1 {
			tmpSum += cost[i]
			if cost[i] > tmpMax {
				tmpMax = cost[i]
			}
			rs += tmpSum - tmpMax
		}
		st, et, tmpMax, tmpSum = -1, -1, 0, 0
	}
	return rs
}

//184 ms	9.3 MB
func minCost3(s string, cost []int) int {
	ret := 0
	cs := []rune(s)
	for i := 0; i < len(cs)-1; i++ {
		for j := i + 1; j < len(cs); j++ {
			if cs[i] == '_' {
				break
			}
			if cs[j] == '_' {
				continue
			}
			if cs[i] != cs[j] {
				break
			}
			if cost[i] >= cost[j] {
				ret += cost[j]
				cs[j] = '_'
			} else {
				ret += cost[i]
				cs[i] = '_'
				break
			}
		}
	}
	return ret
}
