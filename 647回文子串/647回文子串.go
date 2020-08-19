package main

import "fmt"

func main() {
	s := "abccba"
	//s = "aaa"
	fmt.Println(countSubstrings(s))
	fmt.Println(countSubstrings1(s))
	fmt.Println(countSubstringsOfficial(s))
	fmt.Println(countSubstringsOfficial2(s))
}

func countSubstrings(s string) int {
	var count int
	var m = map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if _, ok := m[s[i:j]]; ok {
				count++
				continue
			}
			if isPalindromeString(s[i:j]) {
				m[s[i:j]] = true
				count++
			}
		}
	}
	return count
}

func countSubstrings1(s string) int {
	var count int
	count += len(s)
	var m = map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			if _, ok := m[s[i:j]]; ok {
				count++
				continue
			}
			if isPalindromeString(s[i:j]) {
				m[s[i:j]] = true
				count++
			}
		}
	}
	return count
}

func isPalindromeString(x string) bool {
	if len(x) == 0 {
		//在本例中，空字符串是false，是个例外
		return false
	}
	n := len(x)
	for i := 0; i < n/2; i++ {
		if x[i] != x[n-1-i] {
			return false
		}
	}
	return true
}

func countSubstringsOfficial(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func countSubstringsOfficial2(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}
		// 中心拓展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		// 动态维护 iMax 和 rMax
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整
		ans += f[i] / 2
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
