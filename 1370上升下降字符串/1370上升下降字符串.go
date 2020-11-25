package main

import "fmt"

func main() {
	s := "aaaabbbbcccc"

	fmt.Println(sortString(s))
}

//24 ms	6.9 MB
func sortString0(s string) string {
	var m = make(map[int32]int)
	var minVal, maxVal int32
	minVal, maxVal = 30, 0
	for _, v := range s {
		val := v - 'a'
		m[val]++
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	var ans string
	var sequence bool = true
	for len(m) != 0 {
		if sequence {
			for i := minVal; i <= maxVal; i++ {
				if _, has := m[i]; has {
					ans += string('a' + i)
					m[i]--
					if m[i] == 0 {
						delete(m, i)
					}
				}
			}
			sequence = false
		} else {
			for i := maxVal; i >= minVal; i-- {
				if _, has := m[i]; has {
					ans += string('a' + i)
					m[i]--
					if m[i] == 0 {
						delete(m, i)
					}
				}
			}
			sequence = true
		}
	}

	return ans
}

//4 ms-79.81%	3.2 MB-55.77%
func sortString(s string) string {
	arr := ['z' - 'a' + 1]int{}
	for _, v := range s {
		arr[v-'a']++
	}
	ans := []byte{}

	for len(ans) < len(s) {
		for i := 0; i < 26; i++ {
			if arr[i] != 0 {
				ans = append(ans, byte('a'+i))
				arr[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if arr[i] != 0 {
				ans = append(ans, byte('a'+i))
				arr[i]--
			}
		}
	}
	return string(ans)
}

//0 ms-100.00%	3.2 MB-50%
func sortString1(s string) string {
	n := int('z' - 'a' + 1)
	arr := make([]int, n)
	for _, v := range s {
		arr[v-'a']++
	}
	ans := []byte{}

	for len(ans) < len(s) {
		for i := 0; i < n; i++ {
			if arr[i] != 0 {
				ans = append(ans, byte('a'+i))
				arr[i]--
			}
		}
		for i := n - 1; i >= 0; i-- {
			if arr[i] != 0 {
				ans = append(ans, byte('a'+i))
				arr[i]--
			}
		}
	}
	return string(ans)
}
