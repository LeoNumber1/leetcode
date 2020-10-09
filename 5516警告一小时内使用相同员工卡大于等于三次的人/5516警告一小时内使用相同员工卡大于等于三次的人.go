package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}

	keyName = []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}
	keyTime = []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}

	keyName = []string{"john", "john", "john"}
	keyTime = []string{"23:58", "23:59", "00:01"}

	keyName = []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}
	keyTime = []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}

	fmt.Println(alertNames(keyName, keyTime))
}

//156 ms-100.00%	16.5 MB-100.00%
func alertNames(keyName []string, keyTime []string) []string {
	if len(keyName) != len(keyTime) {
		return nil
	}
	m := map[string][]int{}
	for i := 0; i < len(keyTime); i++ {
		num, _ := strconv.Atoi(strings.Replace(keyTime[i], ":", "", 1))
		m[keyName[i]] = append(m[keyName[i]], num)
	}
	var ans []string
	for k, v := range m {
		n := len(v)
		if n < 3 {
			continue
		}
		sort.Ints(v)
		for i := 0; i < n-1; i++ {
			if i+2 < n && v[i+2]-v[i] <= 100 {
				ans = append(ans, k)
				break
			}
		}
	}
	sort.Strings(ans)
	return ans
}
