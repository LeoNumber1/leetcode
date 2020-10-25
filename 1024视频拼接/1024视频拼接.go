package main

import (
	"fmt"
	"math"
)

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	T := 10

	fmt.Println(videoStitching(clips, T))
}

func videoStitching(clips [][]int, t int) int {
	const inf = math.MaxInt64 - 1
	dp := make([]int, t+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= t; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			// 若能剪出子区间 [l,i]，则可以从 dp[l] 转移到 dp[i]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[t] == inf {
		return -1
	}
	return dp[t]
}
