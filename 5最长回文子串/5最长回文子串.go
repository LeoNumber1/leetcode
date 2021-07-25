package main

import "fmt"

func main() {
	s := "bb"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
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
