package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	candidates = []int{2, 5, 2, 1, 2}
	target = 5

	//fmt.Println(combinationSum2(candidates, target))
	fmt.Println(combinationSum20(candidates, target))
	//fmt.Println(combinationSum21(candidates, target))
	//fmt.Println(combinationSum2Official0(candidates, target))
	fmt.Println(combinationSum2Official1(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	m := map[string]bool{}
	var dfs func(start, sum int, arr []int)
	dfs = func(start, sum int, arr []int) {
		for i := start; i < len(candidates); i++ {
			tmp := sum + candidates[i]
			if tmp < target {
				arr = append(arr, candidates[i])
				dfs(i+1, tmp, arr)
				arr = arr[:len(arr)-1]
			} else if tmp == target {
				t := make([]int, len(arr))
				copy(t, arr)
				t = append(t, candidates[i])
				var s string
				for _, v := range t {
					s += strconv.Itoa(v) + ","
				}
				if _, has := m[s]; !has {
					res = append(res, t)
					m[s] = true
				}
			} else {
				return
			}
		}
	}

	dfs(0, 0, []int{})

	return res
}

func combinationSum20(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	var dfs func(start, sum int, arr []int)
	dfs = func(start, sum int, arr []int) {
		for i := start; i < len(candidates); i++ {
			tmp := sum + candidates[i]

			if tmp > target {
				return
			}

			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if tmp < target {
				arr = append(arr, candidates[i])
				dfs(i+1, tmp, arr)
				arr = arr[:len(arr)-1]
			} else if tmp == target {
				arr = append(arr, candidates[i])
				res = append(res, append([]int(nil), arr...))
				return
			}
		}
	}

	dfs(0, 0, []int{})

	return res
}

//这个是根据【39数组总和】的官方题解改的
func combinationSum2Official0(candidates []int, target int) [][]int {
	comb := []int{}
	res := [][]int{}
	m := map[string]bool{}
	var dfs func(target, index int)
	dfs = func(target, index int) {

		if target == 0 {
			t := append([]int(nil), comb...)
			sort.Ints(t)
			var s string
			for _, v := range t {
				s += strconv.Itoa(v) + ","
			}
			if _, has := m[s]; !has {
				m[s] = true
				res = append(res, t)
			}
			return
		}
		if index == len(candidates) {
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index+1)
			comb = comb[:len(comb)-1]
		}
	}

	dfs(target, 0)

	return res
}

func combinationSum2Official1(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var ans [][]int

func combinationSum21(candidates []int, target int) [][]int {
	ans = [][]int{}
	path := []int{}
	sort.Ints(candidates)
	backtrack(candidates, target, 0, path)
	return ans
}

func backtrack(candidates []int, target, start int, path []int) {
	if target == 0 {
		ans = append(ans, append([]int(nil), path...))
		return
	}

	n := len(candidates)

	for i := start; i < n; i++ {
		if target-candidates[i] < 0 {
			return
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrack(candidates, target-candidates[i], i+1, path)
		path = path[:len(path)-1]
	}
}
