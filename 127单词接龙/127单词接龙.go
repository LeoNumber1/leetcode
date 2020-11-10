package main

import (
	"fmt"
	"math"
)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog", "hat"}

	beginWord = "qa"
	endWord = "sq"
	wordList = []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}

	//fmt.Println(ladderLength0(beginWord, endWord, wordList))
	//fmt.Println(ladderLength(beginWord, endWord, wordList))
	fmt.Println(ladderLengthOfficial(beginWord, endWord, wordList))
}

//时间复杂度超了
func ladderLength0(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	wordLen := len(beginWord)
	if wordLen != len(endWord) || n == 0 {
		return 0
	}
	var has bool
	var isVisit = make([]bool, n)
	for k, v := range wordList {
		if len(v) != wordLen {
			isVisit[k] = true
		}
		if v == endWord {
			has = true
			//break
		}
	}
	if !has {
		return 0
	}

	var ans int
	var dfs func(beginWord string, temp int)
	dfs = func(beginWord string, temp int) {
		if ans != 0 && temp > ans {
			return
		}
		if beginWord == endWord {
			temp++
			if ans == 0 {
				ans = temp
			} else if temp < ans {
				ans = temp
			}
			return
		}
		for k, v := range wordList {
			if !isVisit[k] { //如果没访问过
				for i := 0; i < wordLen; i++ {
					if beginWord[i] != v[i] {
						if i != wordLen-1 { //不是最后一位
							if beginWord[i+1:] == v[i+1:] { //剩下的后面都一样
								isVisit[k] = true
								dfs(v, temp+1)
								isVisit[k] = false
							} else {
								break
							}
						} else { //如果是最后一位
							isVisit[k] = true
							dfs(v, temp+1)
							isVisit[k] = false
						}
					}
				}
			}
		}
	}
	dfs(beginWord, 0)
	return ans
}

func ladderLengthOfficial(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}
