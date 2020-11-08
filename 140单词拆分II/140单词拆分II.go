package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}

	//s = "a"
	//wordDict = []string{"a"}

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	s = "aaaa"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	fmt.Println(wordBreak(s, wordDict))
	fmt.Println(wordBreakOfficial(s, wordDict))
}

//0 ms-100.00%	2.7 MB-48%
//hard 难度，没看答案自己做出来的，撒花*★,°*:.☆(￣▽￣)/$:*.°★* 。
func wordBreak(s string, wordDict []string) []string {
	lens := len(s)
	if lens == 0 || len(wordDict) == 0 {
		return nil
	}
	wordDictMap := map[string]bool{}    //单词字典map，方便查找
	wordMap := map[int32]bool{}         //字母map，递归前检查
	var min, max int = math.MaxInt32, 0 //记录下单词字典的最小、最大长度
	for _, v := range wordDict {
		n := len(v)
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
		for _, val := range v {
			wordMap[val] = true
		}
		wordDictMap[v] = true
	}

	//检查s中是否有字母map里没有的字母
	for _, v := range s {
		if _, has := wordMap[v]; !has {
			return nil
		}
	}

	var ans []string
	var dfs func(strArr []string, index int)
	dfs = func(strArr []string, index int) {
		if index == lens {
			//如果index查完了最后一位，说明找到一组可行解
			ans = append(ans, strings.Join(strArr, " "))
			return
		}
		//index加上从min到max，去wordDictMap查询是否存在
		for i := min; i <= max && i+index <= lens; i++ {
			if _, has := wordDictMap[s[index:index+i]]; has {
				strArr = append(strArr, s[index:index+i])
				dfs(strArr, index+i)
				//剪枝，把strArr中最后一位剪掉
				strArr = strArr[:len(strArr)-1]
			}
		}
	}
	dfs(nil, 0)
	return ans
}

//0 ms-100.00%	2.8 MB-9.82%
func wordBreakOfficial(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
