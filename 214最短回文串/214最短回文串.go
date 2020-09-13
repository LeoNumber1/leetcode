package main

import "fmt"

func main() {
	s := "aacecaaa"

	//s = "ab"
	//s = "abcccd"
	//s = "aab"
	//s = "aba"

	//fmt.Println(shortestPalindrome(s))
	//fmt.Println(shortestPalindromeOfficial(s))
	fmt.Println(shortestPalindromeOfficial2(s))
}

//224 ms	3 MB
func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var subS string
	for i := len(s); i >= 1; i-- {
		if isPalindromeString(s[:i]) {
			subS = s[i:]
			break
		}
	}

	return getPalindromeString(subS) + s
}

func isPalindromeString(x string) bool {
	n := len(x)
	for i := 0; i < n/2; i++ {
		if x[i] != x[n-1-i] {
			return false
		}
	}
	return true
}

func getPalindromeString(x string) string {
	n := len(x)
	if n <= 1 {
		return x
	}
	r := []rune(x)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

//0 ms	2.8 MB
func shortestPalindromeOfficial(s string) string {
	n := len(s)
	base, mod := 131, 1000000007
	left, right, mul := 0, 0, 1
	best := -1
	for i := 0; i < n; i++ {
		left = (left*base + int(s[i]-'0')) % mod
		right = (right + mul*int(s[i]-'0')) % mod
		if left == right {
			best = i
		}
		mul = mul * base % mod
	}
	add := ""
	if best != n-1 {
		add = s[best+1:]
	}
	b := []byte(add)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b) + s
}

//0 ms	3.3 MB
func shortestPalindromeOfficial2(s string) string {
	n := len(s)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && s[j+1] != s[i] {
			j = fail[j]
		}
		if s[j+1] == s[i] {
			fail[i] = j + 1
		}
	}
	best := -1
	for i := n - 1; i >= 0; i-- {
		for best != -1 && s[best+1] != s[i] {
			best = fail[best]
		}
		if s[best+1] == s[i] {
			best++
		}
	}
	add := ""
	if best != n-1 {
		add = s[best+1:]
	}
	b := []byte(add)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b) + s
}
