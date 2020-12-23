package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	//prices = []int{9, 8, 7, 1, 2}
	//fee = 3

	//fmt.Println(maxProfit0(prices, fee))
	fmt.Println(maxProfit(prices, fee))
	fmt.Println(maxProfitOfficial(prices, fee))
}

func maxProfit0(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][2]int, n) //dp[i][j]，i为天数，j为0代表当天不持有，j为1代表当天持有一只股票，值为最大收益
	dp[0][0] = 0            //当前不持有
	dp[0][1] = -prices[0]   //当前持有一只股票，值为最大利润
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], prices[i]+dp[i-1][1]-fee) //上一天不持有和当天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])     //上一天持有和当天买入
	}
	return dp[n-1][0]
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	for i := 1; i < n; i++ {
		dp0 = max(dp0, prices[i]+dp1-fee)
		dp1 = max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitOfficial(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
