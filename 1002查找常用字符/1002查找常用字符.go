package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	A := []string{"bella", "label", "roller"}
	//A = []string{"cool", "lock", "cook"}
	//A = []string{"cool"}

	fmt.Println(commonChars(A))
	//fmt.Println(commonCharsOfficial(A))
}

//4 ms-82.83%	4.8 MB-5.95%
func commonChars1(A []string) []string {
	n := len(A)
	if n == 0 {
		return nil
	} else if n == 1 {
		return strings.Split(A[0], "")
	}
	m := map[int32][]int{}
	for k, v := range A[0] {
		m[v] = append(m[v], k)
	}

	for i := 1; i < n; i++ {
		var tmp = map[int32][]int{}
		for k, v := range A[i] {
			if len(m[v]) > 0 {
				m[v] = m[v][1:]
				tmp[v] = append(tmp[v], k)
			}
		}
		m = tmp
	}
	var ans []string
	for k, v := range m {
		for _, _ = range v {
			ans = append(ans, string(k))
		}
	}
	return ans
}

//8 ms	3.6 MB
func commonChars(A []string) []string {
	n := len(A)
	if n == 0 {
		return nil
	} else if n == 1 {
		return strings.Split(A[0], "")
	}
	m := map[int32]int{}
	for _, v := range A[0] {
		m[v]++
	}

	for i := 1; i < n; i++ {
		var tmp = map[int32]int{}
		for _, v := range A[i] {
			if m[v] > 0 {
				m[v]--
				tmp[v]++
			}
		}
		m = tmp
	}
	var ans []string
	for k, v := range m {
		for i := 0; i < v; i++ {
			ans = append(ans, string(k))
		}
	}
	return ans
}

//0 ms-100%	2.8 MB-100%
func commonCharsOfficial(A []string) (ans []string) {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range A {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}
