package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	s := "123456579"

	fmt.Println(splitIntoFibonacci(s))
	//fmt.Println(splitIntoFibonacciOfficial(s))
}

func splitIntoFibonacciOfficial(s string) (F []int) {
	n := len(s)
	var backtrack func(index, sum, prev int) bool
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}

		cur := 0
		for i := index; i < n; i++ {
			// 每个块的数字一定不要以零开头，除非这个块是数字 0 本身
			if i > index && s[index] == '0' {
				break
			}

			cur = cur*10 + int(s[i]-'0')
			// 拆出的整数要符合 32 位有符号整数类型
			if cur > math.MaxInt32 {
				break
			}

			// F[i] + F[i+1] = F[i+2]
			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}

			// cur 符合要求，加入序列 F
			F = append(F, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			F = F[:len(F)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return
}

func splitIntoFibonacci(S string) []int {
	n1 := 0
	n2 := 0
	for i := 0; i < len(S)/2; i++ {
		if S[0] == '0' && i == 0 {
			n1 = 0
		} else if S[0] == '0' {
			break
		} else {
			n1, _ = strconv.Atoi(S[:i+1])
			if n1 > (math.MaxInt32) {
				break
			}
		}
		for j := i + 1; j < len(S); j++ {
			if S[i+1] == '0' && j == i+1 {
				n2 = 0
			} else if S[i+1] == '0' {
				break
			} else {
				n2, _ = strconv.Atoi(S[i+1 : j+1])
			}
			if n2 > (math.MaxInt32) {
				break
			}
			if len(S[j+1:]) < i || len(S[j+1:]) < j-i {
				break
			}
			//fmt.Println(n1,",",n2,"--",S[j+1:])
			ans, b := doSplitIntoFibonacci(n1, n2, S[j+1:])
			if b {
				return ans
			}
		}
	}
	return make([]int, 0)
}

func doSplitIntoFibonacci(n1, n2 int, s string) ([]int, bool) {
	n := 0
	s1 := ""
	ans := make([]int, 0)
	ans = append(ans, n1)
	ans = append(ans, n2)
	for len(s) >= len(s1) {
		n = n1 + n2
		s1 = strconv.Itoa(n)
		ans = append(ans, n)
		if n > (math.MaxInt32) {
			return nil, false
		}
		if len(s) >= len(s1) && s1 == s[:len(s1)] {
			n1 = n2
			n2 = n
			s = s[len(s1):]
		} else {
			return nil, false
		}
	}
	if s != "" {
		return nil, false
	}
	return ans, true
}
