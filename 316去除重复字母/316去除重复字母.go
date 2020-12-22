package main

import "fmt"

func main() {
	s := "cbacdcbc"
	//s = "bcabc"

	fmt.Println(removeDuplicateLetters(s))
	//fmt.Println(removeDuplicateLettersOfficial(s))
}

func removeDuplicateLetters(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	m := make(map[int32]int, 26)
	for _, ch := range s {
		m[ch]++
	}
	stack := []int32{}
	inStack := [26]bool{}
	for _, ch := range s {
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				//当栈不为空，且当前字符比栈顶元素小，则检查栈顶的元素后面还有没，如果还有，则把栈里的大于当前字符的元素清理掉
				top := stack[len(stack)-1]
				if m[top] > 0 {
					stack = stack[:len(stack)-1]
					inStack[top-'a'] = false
				} else {
					break
				}
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		m[ch]--
	}

	return string(stack)
}

func removeDuplicateLettersOfficial(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
