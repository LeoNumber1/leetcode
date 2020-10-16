package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{-4, -1, 0, 3, 10}

	//fmt.Println(sortedSquares(A))
	//fmt.Println(sortedSquaresOfficial1(A))
	fmt.Println(sortedSquaresOfficial2(A))
}

//40 ms	6.8 MB
func sortedSquares(A []int) []int {
	for k, v := range A {
		A[k] = v * v
	}
	sort.Ints(A)
	return A
}

//36 ms	6.8 MB
func sortedSquaresOfficial1(A []int) []int {
	n := len(A)
	lastNegIndex := -1
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			lastNegIndex = i
		} else {
			break
		}
	}
	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, A[j]*A[j])
			j++
		} else if j == n {
			ans = append(ans, A[i]*A[i])
			i--
		} else if A[i]*A[i] < A[j]*A[j] {
			ans = append(ans, A[i]*A[i])
			i--
		} else {
			ans = append(ans, A[j]*A[j])
			j++
		}
	}
	return ans
}

//36 ms-77.52%	6.6 MB-65.75%
func sortedSquaresOfficial2(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := a[i]*a[i], a[j]*a[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}
