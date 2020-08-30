package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"

	fmt.Println(reverseWords1(s))
	fmt.Println(reverseWords2(s))
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for k, v := range arr {
		arr[k] = reverse(v)
	}
	return strings.Join(arr, " ")
}

func reverse(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		res += string(s[len(s)-1-i])
	}
	return res
}

func reverseWords1(s string) string {
	arr := strings.Split(s, " ")
	for k, v := range arr {
		arr[k] = reverse1(v)
	}
	return strings.Join(arr, " ")
}

func reverse1(s string) string {
	arr := strings.Split(s, "")
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, "")
}

func reverseWordsOfficial(s string) string {
	length := len(s)
	ret := []byte{}
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start+i-1-p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

func reverseWords2(s string) string {
	b := []byte(s)
	l := 0
	for k, v := range s {
		if v == ' ' || k == len(s)-1 {
			r := k - 1
			if k == len(s)-1 {
				r = k
			}
			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}
			l = k + 1
		}
	}
	return string(b)
}
