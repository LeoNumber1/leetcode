package main

import (
	"fmt"
	"time"
)

func main() {
	haystack := "hello"
	needle := "ll"

	haystack = "baaabba"
	needle = "bba"

	//haystack = "a"
	//needle = "a"

	haystack = "mississippi"
	needle = "issip"

	t0 := time.Now()
	fmt.Println(strStr(haystack, needle), time.Since(t0))
	t1 := time.Now()
	fmt.Println(strStrOfficial1(haystack, needle), time.Since(t1))
	t2 := time.Now()
	fmt.Println(strStrKMP(haystack, needle), time.Since(t2))
}

func strStr(haystack string, needle string) int {
	if needle == "" { //needle为空返回0
		return 0
	}
	if len(needle) > len(haystack) {
		return -1 //needle长度大于haystack返回-1
	}
	var i, j int   //定义指针i，j分别对应haystack和needle
	var index = -1 //index位置初始值-1
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if index == -1 {
				index = i //index是初始值-1，则将当前i记录下来
			}
			if j == len(needle)-1 {
				return index //如果当前的j已经是needle的最后一位，则返回index
			}
			j++ //否则needle的指针j往后移动一位
		} else {
			if index != -1 {
				i = index //如果index不为初始值-1，则将i赋值为index，下次从index的下一个位置开始继续找
				index = -1
				j = 0 //再将index和j赋值回初始值
			}
		}
		i++ //不管上面怎么变动，haystack的指针每次都要往后移动一位
	}

	return -1
}

func strStrOfficial1(haystack string, needle string) int {
	if needle == "" { //needle为空返回0
		return 0
	}
	if len(needle) > len(haystack) {
		return -1 //needle长度大于haystack返回-1
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func KMPNext(s string) []int {
	lenth := len(s)
	next := make([]int, lenth)
	next[0] = 0
	i, j := 1, 0
	for i < lenth {
		if s[i] == s[j] {
			next[i] = j + 1 // 一下个匹配位置为下一位
			i++
			j++
		} else {
			if j == 0 {
				next[i] = 0 // 重头开始匹配
				i++
			} else {
				j = next[j-1] // 回退
			}
		}
	}
	return next
}

func strStrKMP(haystack string, needle string) int {
	lenRoot := len(haystack)
	lenTmpl := len(needle)
	if lenTmpl == 0 {
		return 0
	}

	next := KMPNext(needle)
	for i, j := 0, 0; i < lenRoot; {
		for j < lenTmpl && i < lenRoot && haystack[i] == needle[j] {
			i++
			j++
		}
		if j == lenTmpl {
			return i - j
		}
		if i == lenRoot {
			return -1
		}
		if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
	}
	return -1
}
