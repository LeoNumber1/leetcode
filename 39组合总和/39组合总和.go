package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	candidates := []int{
		//2, 3, 6, 7,
		2, 3, 5,
	}
	target := 7
	//target = 8

	t := time.Now()
	fmt.Println(combinationSum(candidates, target), time.Since(t))
	to := time.Now()
	fmt.Println(combinationSumOfficial(candidates, target), time.Since(to))
	t1 := time.Now()
	fmt.Println(combinationSum1(candidates, target), time.Since(t1))
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

var result [][]int
var curSel []int

func combinationSum1(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = make([][]int, 0)
	curSel = make([]int, 0)
	DFS(target, candidates)
	return result
}

func DFS(target int, candidates []int) int {
	if getSum(curSel) == target {
		cCurSel := make([]int, len(curSel))
		copy(cCurSel, curSel)
		result = append(result, cCurSel)
		return 0
	} else if getSum(curSel) > target {
		return -1
	} else { ////主要看这里用0代表相同，-1代表已经超过了当前target，1则表示还能继续加
		for i := 0; i < len(candidates); i++ {
			curSel = append(curSel, candidates[i])
			temp := DFS(target, candidates[i:])
			curSel = curSel[:len(curSel)-1]
			if temp <= 0 {
				break
			}
		}
	}
	return 1
}

func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
