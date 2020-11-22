package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "anagram"
	t := "nagaram"

	//s = "rat"
	//t = "car"

	fmt.Println(isAnagram(s, t))
}

//8 ms	2.8 MB
func isAnagram0(s string, t string) bool {
	var m = make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		if _, has := m[v]; has {
			if m[v] > 1 {
				m[v]--
			} else {
				delete(m, v)
			}
		} else {
			return false
		}
	}
	return len(m) == 0
}

//12 ms	3 MB
func isAnagram(s string, t string) bool {
	s1, t1 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	return string(s1) == string(t1)
}

func isAnagramOfficial(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1, c2 [26]int
	for k, v := range s {
		c1[v-'a']++
		c2[t[k]-'a']++
	}
	return c1 == c2
}
