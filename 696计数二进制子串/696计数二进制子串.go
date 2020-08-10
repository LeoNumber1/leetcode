package main

import "fmt"

func main() {
	s := "00110011"
	//s = "10101"
	//fmt.Println(countBinarySubstrings(s))
	fmt.Println(countBinarySubstringsOfficial(s))
	fmt.Println(countBinarySubstringsOfficial1(s))
}

func countBinarySubstrings(s string) int {
	var count, i int
	for j := i + 1; i < len(s) && j < len(s); i, j = i+1, j+1 {
		if s[i] != s[j] {
			count++
			for x, y := i-1, j+1; x >= 0 && y < len(s); x, y = x-1, y+1 {
				if s[x] == s[i] && s[y] == s[j] {
					count++
				} else {
					break
				}
			}
		}
	}
	return count
}

//时间复杂度和空间复杂度都是 O(n)O(n)。
func countBinarySubstringsOfficial(s string) int {
	counts := []int{}
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
	}
	var ans int
	for i := 1; i < len(counts); i++ {
		ans += min(counts[i], counts[i-1])
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//时间复杂度是 O(n),空间复杂度O(1)。
func countBinarySubstringsOfficial1(s string) int {
	var ptr, last, ans int
	n := len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		ans += min(count, last)
		last = count
	}

	return ans
}
