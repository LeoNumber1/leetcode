package main

import "fmt"

func main() {
	s := "abcabcbb" //3
	s = "bbbbb"
	s = "pwwkew"   //3
	s = "pwekwkew" //4
	s = "au"       //2
	s = " "
	s = ""
	s = "aadc er"  //6
	s = "abba"     //2
	s = "tmmzuxt"  //5
	s = "bbtablud" //6

	//fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstringOfficial(s))
}

//12 ms	3.1 MB
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	m := map[byte]int{} //value值为下标值+1
	var lent, maxLenth int
	index := 0 //index为当前不重复字符串开始的下标
	for i := 0; i < n; i++ {
		if v, has := m[s[i]]; !has || (has && v < index) {
			lent++
			m[s[i]] = i + 1
		} else {
			//进这里的v是重复字符串的位置+1
			if lent > maxLenth {
				maxLenth = lent
			}
			m[s[i]] = i + 1
			lent = i - v + 1
			index = v
		}
	}
	if lent > maxLenth {
		return lent
	}
	return maxLenth
}

//12 ms	2.8 MB
func lengthOfLongestSubstringOfficial(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
		if rk+1 == n {
			break
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
