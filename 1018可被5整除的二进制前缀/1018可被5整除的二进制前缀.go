package main

import (
	"fmt"
)

func main() {
	A := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}

	fmt.Println(prefixesDivBy5(A))
}

func prefixesDivBy5(A []int) []bool {
	ans := make([]bool, len(A))
	pre := 0
	for k, v := range A {
		pre = (pre<<1 | v) % 5
		ans[k] = pre == 0
	}
	return ans
}
