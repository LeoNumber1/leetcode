package main

import "fmt"

func main() {
	s := "abab"
	s = "aba"
	s = "abcabcabcabc"
	s = "ababab"
	s = "babbabbabbabbab"
	//s = "aabaaba"

	//fmt.Println(repeatedSubstringPattern(s))
	//fmt.Println(repeatedSubstringPatternOfficial(s))
	fmt.Println(repeatedSubstringPatternOfficial1(s))
}

func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	}

	var tmp string
	for i := 0; i < len(s)/2; i++ {
		tmp += string(s[i])
		for j := i + 1; j < len(s); j = j + len(tmp) {
			//fmt.Println("tmp =", tmp)
			//fmt.Println("当前串：", s[j:j+len(tmp)])
			if j+len(tmp) <= len(s) && tmp != s[j:j+len(tmp)] {
				break
			} else {
				if j+len(tmp) == len(s) {
					return true
				}
			}
		}
	}
	return false
}

func repeatedSubstringPatternOfficial(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPatternOfficial1(s string) bool {
	return kmp(s+s, s)
}

func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}
	for i := 1; i < m; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	match := -1
	for i := 1; i < n-1; i++ {
		for match != -1 && pattern[match+1] != query[i] {
			match = fail[match]
		}
		if pattern[match+1] == query[i] {
			match++
			if match == m-1 {
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPatternOfficial2(s string) bool {
	return kmp1(s)
}

func kmp1(pattern string) bool {
	n := len(pattern)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	return fail[n-1] != -1 && n%(n-fail[n-1]-1) == 0
}
