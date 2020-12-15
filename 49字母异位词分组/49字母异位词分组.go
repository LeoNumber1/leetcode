package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	//fmt.Println(groupAnagrams(strs))
	//fmt.Println(groupAnagramsOfficial(strs))
	fmt.Println(groupAnagramsOfficial1(strs))
}

//32 ms-50.89%	7.9 MB-56.21%
func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	if n == 0 {
		return nil
	} else if n == 1 {
		return [][]string{strs}
	}
	type myStruct struct {
		str    []byte
		index  int
		length int
	}
	myStructs := []myStruct{}
	for k, str := range strs {
		temp := []byte(str)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		myStructs = append(myStructs, myStruct{temp, k, len(str)})
	}
	sort.Slice(myStructs, func(i, j int) bool {
		return myStructs[i].length < myStructs[j].length
	})
	ans := [][]string{}
	start := 0
	count := n
	for count > 0 {
		length := myStructs[start].length
		myMap := make(map[string][]int)
		for i := start; i < n; i++ {
			start = i
			if myStructs[i].length != length {
				length = myStructs[i].length
				break
			}
			myMap[string(myStructs[i].str)] = append(myMap[string(myStructs[i].str)], myStructs[i].index)
		}
		for _, ints := range myMap {
			temp := []string{}
			for _, v := range ints {
				temp = append(temp, strs[v])
				count--
			}
			ans = append(ans, temp)
		}
	}
	return ans
}

//32 ms-50.89%	7.7 MB-64.85%
func groupAnagramsOfficial(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagramsOfficial1(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
