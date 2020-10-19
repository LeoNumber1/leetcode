package main

import (
	"fmt"
)

func main() {
	S := "ab#c"
	T := "ad#c"

	S = "ab##"
	T = "c#d#"

	S = "a#c"
	T = "b"

	//S = "a##c"
	//T = "#a#c"

	//fmt.Println(backspaceCompare(S, T))
	fmt.Println(backspaceCompareOfficial(S, T))
}

//0 ms-100.00%	2 MB-53.85%
func backspaceCompare(S string, T string) bool {
	return do(S) == do(T)
}

func do(s string) string {
	var arr []byte
	for _, v := range s {
		if v == '#' {
			if len(arr) > 0 {
				arr = arr[:len(arr)-1]
			}
		} else {
			arr = append(arr, byte(v))
		}
	}
	return string(arr)
}

//0 ms-100.00%	2 MB-67.03%
func backspaceCompareOfficial(s string, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
