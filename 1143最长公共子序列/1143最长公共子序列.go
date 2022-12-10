package main

import "fmt"

func main() {
	s1 := "abasdssad"
	s2 := "acsassdfag"

	fmt.Println(longestCommonSubsequence(s1, s2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	dp := make([]int, m+1)
	for _, c2 := range text2 {
		var last int
		for i, c1 := range text1 {
			tmp := dp[i+1]
			if c1 == c2 {
				dp[i+1] = last + 1
			} else {
				dp[i+1] = max(tmp, dp[i])
			}
			last = tmp
		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
