package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}
	//cost = []int{1, 2}
	//cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	//fmt.Println(minCostClimbingStairs0(cost))
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs0(cost []int) int {
	n := len(cost)
	if n == 2 {
		return min(cost[0], cost[1])
	}
	dp := make([][2]int, n) //dp[i][j] i为楼梯索引，j=0为不踩这阶楼梯，j=1为踩这阶楼梯，所花费的总体力值
	dp[0][0] = 0
	dp[0][1] = cost[0]
	dp[1] = [2]int{0, cost[1]}
	for i := 2; i < n; i++ {
		dp[i][0] = dp[i-1][1]                            //这阶不踩，上一阶必须踩
		dp[i][1] = min(dp[i-1][1], dp[i-2][1]) + cost[i] //这阶踩，取上一阶和上上阶踩的最小值，
	}
	return min(dp[n-1][0], dp[n-1][1])
}

func minCostClimbingStairs(cost []int) int {
	dp0, dp1 := cost[0], [2]int{cost[0], cost[1]}
	for i := 2; i < len(cost); i++ {
		dpNew := [2]int{dp1[1], min(dp0, dp1[1]) + cost[i]}
		dp0, dp1 = dp1[1], dpNew
	}
	return min(dp1[0], dp1[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
