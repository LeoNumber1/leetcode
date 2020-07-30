package main

import "fmt"

func main() {
	strs := []string{
		"flower", "flow", "flight",
		//"dog", "racecar", "car",
		//"abab", "aba", "",
	}
	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix1(strs))
	fmt.Println(longestCommonPrefixOfficial1(strs))
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	var minLenth int = len(strs[0])
	//arr := [][]string{{}}
	for i := 0; i < n; i++ {
		minLenth = min(minLenth, len(strs[i]))
	}
	if minLenth == 0 {
		return ""
	}
	var res string = strs[0][:minLenth]
for1:
	for j := 0; j < minLenth; j++ {
		for i := 1; i < n; i++ {
			if j < len(res) {
				if res[j] != strs[i][j] {
					res = res[:j]
					break for1
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestCommonPrefix1(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	var res string = strs[0]
	for i := 1; i < n; i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		for j := 0; j < len(strs[i]); j++ {
			if j < len(res) {
				if j == len(strs[i])-1 {
					if strs[i][j] != res[j] {
						res = res[:j]
					} else {
						res = res[:j+1]
					}
				} else {
					if strs[i][j] != res[j] {
						res = res[:j]
					}
				}
			}
		}
	}
	return res
}

func longestCommonPrefixOfficial1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}
