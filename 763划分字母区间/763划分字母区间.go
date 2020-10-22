package main

import (
	"fmt"
	"math"
)

func main() {
	S := "ababcbacadefegdehijhklij"
	//S := "abcabcbacadefwxyz"

	//fmt.Println(partitionLabels(S))
	fmt.Println(partitionLabelsOfficial(S))
}

//	4 ms	2.2 MB
func partitionLabels(S string) []int {
	var ans []int
	var f func(S string)
	f = func(S string) {
		if len(S) == 0 {
			return
		}
		end := 0
		pre := S[0] //待比较字符
		arr := 0    //最短字符串内已经包含的字符下标位置，位表示
		tmp := 0    //最短字符串剩余的字符串里包含的字符下标位置，位表示
	next:
		for i := end + 1; i < len(S); i++ {
			if S[i] == pre {
				//如果当前字符等于待比较字符，则把tmp里的下标全合并到arr中，同时将tmp归零
				arr |= tmp
				tmp = 0
				end = i
			} else {
				//如果当前字符不等于待比较字符，则把该字符的下标计入tmp中
				tmp |= 1 << int(S[i]-'a')
			}
		}
		at := arr & tmp //检查剩下的字符里有没有与和之前重复的
		if at == 0 {
			//没有就记录下来长度
			ans = append(ans, end+1)
		} else {
			//有重复
			tmp = 0
			//最小的重复字符位置
			index := at & -at
			pre = byte('a' + int(math.Log2(float64(index))))
			//跳转回去继续找重复
			goto next
		}
		f(S[end+1:])
	}
	f(S)
	return ans
}

//	0 ms	2.1 MB
func partitionLabelsOfficial(S string) []int {
	var ans []int
	lastPos := [26]int{}
	for k, v := range S {
		lastPos[v-'a'] = k
	}
	start, end := 0, 0
	for k, v := range S {
		if lastPos[v-'a'] > end {
			end = lastPos[v-'a']
		}
		if k == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
