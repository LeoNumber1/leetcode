package main

import (
	"fmt"
	"time"
)

func main() {
	digits := "2345678765439"

	t := time.Now()
	letterCombinations(digits)
	fmt.Println("我的 ：", time.Since(t))

	to1 := time.Now()
	letterCombinationsOfficial(digits)
	fmt.Println("官方解法1 ：", time.Since(to1))
}

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	} else if len(digits) == 1 {
		return m[digits[0]]
	}

	return dfs(digits, 0)
}

func dfs(digits string, start int) []string {
	if start == len(digits)-1 {
		return m[digits[start]]
	}
	var res []string

	for _, v1 := range m[digits[start]] {
		for _, v2 := range dfs(digits, start+1) {
			res = append(res, v1+v2)
		}
	}
	return res
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinationsOfficial(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
