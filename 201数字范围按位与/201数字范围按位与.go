package main

import (
	"fmt"
	"time"
)

func main() {
	m := 5
	n := 7

	//m = 0
	//n = 1

	m = 4000000
	n = 2147483646

	//m = 0
	//n = 4

	t1 := time.Now()
	fmt.Println(rangeBitwiseAnd(m, n), time.Since(t1))
	t2 := time.Now()
	fmt.Println(rangeBitwiseAnd1(m, n), time.Since(t2))
}

func rangeBitwiseAnd(m int, n int) int {
	var res = m
	for i := m + 1; i <= n; i++ {
		res &= i
		if res == 0 {
			return res
		}
	}
	return res
}

func rangeBitwiseAnd1(m int, n int) int {
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
