package main

import (
	"fmt"
	"time"
)

func main() {
	candidates := []int{
		2, 3, 6, 7,
		//2, 3, 5,
	}
	target := 7
	//target = 8

	//t := time.Now()
	//fmt.Println(combinationSum(candidates, target), time.Since(t))
	to := time.Now()
	fmt.Println(combinationSumOfficial(candidates, target), time.Since(to))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}

	var dfs func(start int, tmp []int, sum int)
	dfs = func(start int, tmp []int, sum int) {
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] < target {
				tmp = append(tmp, candidates[i])
				dfs(i, tmp, sum+candidates[i])
				tmp = tmp[:len(tmp)-1]
			} else if sum+candidates[i] == target {
				t := make([]int, len(tmp))
				copy(t, tmp)
				t = append(t, candidates[i])
				res = append(res, t)
			}
		}
	}

	dfs(0, []int{}, 0)

	return res
}

func combinationSumOfficial(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
