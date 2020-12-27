package main

import "fmt"

func main() {
	s := "egg"
	t := "add"

	//fmt.Println(isIsomorphic(s, t))
	fmt.Println(isIsomorphicOfficial(s, t))
}

func isIsomorphic(s string, t string) bool {
	ms := make(map[int32]struct{})
	mt := make(map[int32]struct{})
	for _, ch := range s {
		ms[ch] = struct{}{}
	}
	for _, ch := range t {
		mt[ch] = struct{}{}
	}
	if len(ms) != len(mt) {
		return false
	}
	m := make(map[int32]int32)
	for k, ch := range s {
		if val, has := m[ch]; has {
			if val != int32(t[k]) {
				return false
			}
		} else {
			m[ch] = int32(t[k])
		}
	}
	return true
}

func isIsomorphicOfficial(s, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
