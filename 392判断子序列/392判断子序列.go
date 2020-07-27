package main

import "fmt"

func main() {
	s := "abc"
	s = "axc"
	//s = "exc"
	t := "ahbgdc"
	//fmt.Println(isSubsequence(s, t))
	//fmt.Println(isSubsequenceOfficial1(s, t))
	fmt.Println(isSubsequenceOfficial2(s, t))
}

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls > lt {
		return false
	}
	if ls == 0 {
		return true
	}

	var index int
	for i := 0; i < ls; i++ {
		for j := index; j < lt; j++ {
			if s[i] == t[j] {
				index = j + 1
				if i == ls-1 {
					return true
				}
				break
			}
			if ls-i >= lt-j {
				return false
			}
		}
	}
	return false
}

func isSubsequenceOfficial1(s string, t string) bool {
	n, m := len(s), len(t)
	var i, j int
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

//官方题解，动态规划
func isSubsequenceOfficial2(s string, t string) bool {
	n, m := len(s), len(t)
	f := make([][26]int, m+1)
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}
	add := 0
	for i := 0; i < n; i++ {
		if f[add][int(s[i]-'a')] == m {
			return false
		}
		add = f[add][int(s[i]-'a')] + 1
	}
	return true
}
