package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 4
	k := 9

	//n = 3
	//k = 3

	//for i := 1; i <= k; i++ {
	//	fmt.Println(getPermutation(n, i))
	//}
	fmt.Println(getPermutation(n, k))
	fmt.Println(getPermutationOfficial(n, k))
}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}

	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	factorialArr := factorial(n - 1)

	var s string
	var fun func(n, k int)

	fun = func(n, k int) {
		f := factorialArr[n-1]
		sh := k / f
		yu := k % f
		var index int
		if yu == 0 {
			index = sh - 1
			yu = k - index*f
		} else {
			index = sh
		}
		s += strconv.Itoa(arr[index])
		arr = append(arr[:index], arr[index+1:]...)
		if len(arr) == 1 {
			s += strconv.Itoa(arr[0])
			return
		}
		fun(n-1, yu)
	}

	fun(n, k)
	return s
}

func factorial(n int) []int {
	factorialArr := []int{
		1,
	}
	facVal := 1
	for i := 1; i <= n; i++ {
		facVal *= i
		factorialArr = append(factorialArr, facVal)
	}
	return factorialArr
}

func getPermutationOfficial(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans
}
