package main

import "fmt"

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	f[2] = 2
	//f[3] = 3
	for i := 2; i < n; i++ {
		f[i+1] = f[i] + f[i-1]
	}

	return f[n]
}
