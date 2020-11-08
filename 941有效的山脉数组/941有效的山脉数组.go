package main

import "fmt"

func main() {
	A := []int{0, 3, 2, 1}
	A = []int{3, 2, 1}

	fmt.Println(validMountainArray(A))
}

//28 ms-90.79%	6.4 MB-39.47%
func validMountainArray0(A []int) bool {
	n := len(A)
	i := 0
	for i < n-1 && A[i+1] > A[i] {
		i++
	}
	if i == 0 || i == n-1 {
		return false
	}
	for i < n-1 && A[i+1] < A[i] {
		i++
	}
	return i == n-1
}

func validMountainArray(A []int) bool {
	n := len(A)
	left, right := 0, n-1
	for left < n-1 && A[left] < A[left+1] {
		left++
	}
	for right > 0 && A[right-1] > A[right] {
		right--
	}
	return left > 0 && right < n-1 && left == right
}
