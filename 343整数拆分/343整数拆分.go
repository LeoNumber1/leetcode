package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	n := 10
	//n = 102
	//n = 2
	//n = 9
	t1 := time.Now()
	fmt.Println(integerBreak(n), "数学法——耗时：", time.Since(t1))
	t2 := time.Now()
	fmt.Println(integerBreak1(n), "递归法——耗时：", time.Since(t2))
	t3 := time.Now()
	fmt.Println(integerBreak2(n), "动态规划——耗时：", time.Since(t3))

	for i := 2; i < 59; i++ {
		fmt.Println("i =", i, "————> 值 =", integerBreak(i))
		fmt.Println("i =", i, "————> 值 =", integerBreak1(i))
		fmt.Println("i =", i, "————> 值 =", integerBreak2(i))
	}
}

func integerBreak(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	}

	x := n / 3
	y := n % 3
	switch y {
	case 0:
		return int(math.Pow(3, float64(x)))
	case 1:
		return int(math.Pow(3, float64(x-1))) * 4
	case 2:
		return int(math.Pow(3, float64(x)) * 2)
	}
	return 0
}

var m = map[int]int{}

//递归
func integerBreak1(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	}

	m = make(map[int]int)
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 6

	if n <= 5 {
		return m[n]
	}

	x1 := n / 2
	x2 := n - x1
	return max(IntB(x1)*IntB(x2), IntB(x1-1)*IntB(x2+1))
}

func IntB(n int) int {
	if v, ok := m[n]; ok {
		return v
	}
	x1 := n / 2
	x2 := n - x1
	return max(IntB(x1)*IntB(x2), IntB(x1-1)*IntB(x2+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//动态规划
func integerBreak2(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	}

	m := make(map[int]int)
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 6

	if n <= 5 {
		return m[n]
	}

	for i := 6; i < n+1; i++ {
		x1 := i / 2
		x2 := i - x1
		m[i] = max(m[x1]*m[x2], m[x1-1]*m[x2+1])
	}

	return m[n]
}
