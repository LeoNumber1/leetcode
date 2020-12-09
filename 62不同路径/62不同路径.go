package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePathsOfficial(3, 2))
}

//滚动数组dp    0 ms-100.00%	1.9 MB-96.67%
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if m < n {
		//保证m是较大值
		m, n = n, m
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func uniquePathsOfficial(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
