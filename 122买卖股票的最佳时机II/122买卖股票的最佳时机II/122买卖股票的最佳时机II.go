package main

import "fmt"

func main() {
	price := []int{7, 1, 5, 3, 6, 4}
	price = []int{1, 2, 3, 4, 5}
	price = []int{7, 6, 4, 3, 1}
	price = []int{1, 2}

	fmt.Println(maxProfit(price))
}

//4 ms-95.25%	3 MB-76.89%
func maxProfit0(prices []int) int {
	var ans, buyPrice int = 0, -1
	var has bool
	if len(prices) < 2 {
		return 0
	}

	for i := 0; i < len(prices)-1; i++ {
		if !has {
			if prices[i] < prices[i+1] {
				buyPrice = prices[i]
				has = true
			}
		} else {
			if prices[i] > buyPrice && prices[i+1] <= prices[i] {
				has = false
				ans += prices[i] - buyPrice
				continue
			}
			if has && i+1 == len(prices)-1 {
				ans += prices[i+1] - buyPrice
				has = false
			}
		}
	}
	if has && prices[len(prices)-1]-buyPrice > 0 {
		ans += prices[len(prices)-1] - buyPrice
	}
	return ans
}

//4 ms-95.25%	3 MB-76.89%
func maxProfit1(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n) //i为第几天，j为0代表当天不持有，j=1表示当天持有一只股票，值为最大利润
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
