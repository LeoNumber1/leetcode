package main

import (
	"fmt"
)

func main() {
	s := "cbbd"
	//s = "ccc"
	//s = "cc"
	s = "babad"
	fmt.Println(longestPalindrome(s))
}

//8 ms  2.6 MB      中心扩展法
func longestPalindrome1(s string) string {
	var ans = string(s[0])
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			w := i
			for j := i - 1; j >= 0 && w < n; j-- {
				if s[j] == s[w] {
					str := s[j : w+1]
					if len(str) > len(ans) {
						ans = str
					}
				} else {
					break
				}
				w++
			}
		}
		if i+1 < n && s[i+1] == s[i-1] {
			w := i + 1
			for j := i - 1; j >= 0 && w < n; j-- {
				if s[j] == s[w] {
					str := s[j : w+1]
					if len(str) > len(ans) {
						ans = str
					}
				} else {
					break
				}
				w++
			}
		}
	}
	return ans
}

//136 ms    2.7 MB      暴力法
func longestPalindrome0(s string) string {
	var ans string
	n := len(s)
	if n == 1 {
		return s
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			str := s[i:j]
			if len(str) <= len(ans) {
				continue
			}
			if isPalindromeString(str) {
				if len(str) > len(ans) {
					ans = str
				}
			}
		}
	}
	return ans
}

func isPalindromeString(x string) bool {
	if len(x) == 0 {
		return true
	}
	n := len(x)
	for i := 0; i < n/2; i++ {
		if x[i] != x[n-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen, begin := 1, 0
	// dp[i][j] 表示 s[i...j]是否是回文串
	dp := make([][]bool, n)
	// 初始化：所有长度为 1 的子串都是回文串
	for k := range dp {
		dp[k] = make([]bool, n)
		dp[k][k] = true
	}
	// 递推开始
	// 先枚举子串长度
	for L := 1; L < n; L++ {
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < n; i++ {
			// 由 L 和 i 可以确定右边界，即 j - i = L 得
			j := L + i
			// 如果右边界越界，就可以退出当前循环
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			//	只要 dp[i][L] == true 成立，就表示子串 s[i...L]是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
