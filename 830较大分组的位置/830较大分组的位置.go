package main

import "fmt"

func main() {
	s := "aaa"
	//s = "abbxxxxzzy"

	fmt.Println(largeGroupPositions(s))
	//fmt.Println(largeGroupPositionsOfficial(s))
}

func largeGroupPositions(s string) [][]int {
	n := len(s)
	if n < 3 {
		return nil
	}
	var ans [][]int
	pre := s[0]
	prevKey := 0
	temp := []int{}
	for k := 1; k < n; k++ {
		if s[k] != pre {
			if len(temp) == 2 {
				ans = append(ans, temp)
				temp = []int{}
			}
			pre = s[k]
			prevKey = k
			continue
		}
		if k+1 < n {
			if s[k+1] == pre {
				if len(temp) == 2 {
					temp[1] = k + 1
				} else {
					temp = append(temp, []int{prevKey, k + 1}...)
				}
			}
		}
	}
	if len(temp) == 2 {
		ans = append(ans, temp)
	}
	return ans
}

func largeGroupPositionsOfficial(s string) (ans [][]int) {
	cnt := 1
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				ans = append(ans, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	return
}
