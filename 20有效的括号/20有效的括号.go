package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "()[]{}"
	s = ""
	//s = "()[]"
	//s = "()"
	//s = "{()}"
	//s = "(]"
	//s = "({)}"
	//s = "([(({["
	//fmt.Println(isValid(s))
	//fmt.Println(isValid1(s))
	fmt.Println(isValid2(s))
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	arr := make([]string, 0)
	for i := 0; i < n; i++ {
		if len(arr) > len(s)/2 {
			return false
		}

		if _, ok := m[string(s[i])]; ok {
			arr = append(arr, string(s[i]))
		} else {
			if len(arr) > 0 {
				if string(s[i]) != m[arr[len(arr)-1]] {
					return false
				} else {
					arr = arr[:len(arr)-1]
				}
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}

func isValid1(s string) bool {
	for strings.Contains(s, "()") || strings.Contains(s, "{}") || strings.Contains(s, "[]") {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
	}
	return s == ""
}

func equal(s1 byte, s2 byte) bool {
	sum := s1 + s2
	return sum == '{'+'}' || sum == '('+')' || sum == '['+']'
}

func isValid2(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}

	stack := make([]byte, len(s))
	stackTop := -1

	for i := 0; i < len(s); i++ {
		if stackTop == -1 {
			stackTop++
			stack[stackTop] = s[i]
		} else if equal(stack[stackTop], s[i]) {
			stackTop--
		} else {
			stackTop++
			stack[stackTop] = s[i]
		}
	}

	return stackTop == -1
}
