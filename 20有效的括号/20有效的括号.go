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
	if n%2 == 1 { //如果是奇数，返回false
		return false
	}
	m := map[string]string{ //定义括号对的map
		"(": ")",
		"[": "]",
		"{": "}",
	}

	arr := make([]string, 0)
	for i := 0; i < n; i++ {
		if len(arr) > len(s)/2 { //如果已经遍历了一半，arr还不为空，则直接false
			return false
		}

		if _, ok := m[string(s[i])]; ok { //如果是左括号，则加入arr
			arr = append(arr, string(s[i]))
		} else { //不是左括号就去arr里的最后一个匹配，匹配成功了就删去arr里的最后一个，匹配不成功了直接返回false
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
