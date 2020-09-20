package main

import (
	"fmt"
)

func main() {
	k := 1
	n := 1

	k = 1
	n = 5

	k = 5
	n = 50
	fmt.Println(keyboard1(k, n))
}

const (
	NUM  = 1000000007
	WORD = 26
	//WORD = 3
)

func keyboard(k int, n int) int {
	arr := make([][]int, k)
	for i := 0; i < k; i++ {
		arr[i] = make([]int, n)
	}

	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				arr[i][j] = 26
			} else if j <= i {
				arr[i][j] = arr[i][j-1] * 26 % NUM
			} else {
				arr[i][j] = arr[i][j-1] * (26 - (j - i)) % NUM
			}
		}
	}

	return arr[k-1][n-1] % NUM
}

func keyboard1(k int, n int) int {
	var pre int = WORD
	for i := 1; i < n; i++ {
		if i > WORD-1 {
			break
		}
		if i <= k-1 {
			pre *= WORD
		} else {
			pre *= WORD - (i + 1 - k)
		}
		pre %= NUM
	}

	return pre
}
