package main

import "fmt"

func main() {
	n := 4
	k := 3

	//fmt.Println(combine(n, k))
	//fmt.Println(combineOfficial(n, k))
	fmt.Println(combineOfficial2(n, k))
}

func combine(n int, k int) [][]int {
	res := [][]int{}
	var f func(start int, tmp []int)
	f = func(start int, tmp []int) {
		for i := start; i <= n; i++ {
			tmp = append(tmp, i)
			if len(tmp) == k {
				t := make([]int, len(tmp))
				copy(t, tmp)
				res = append(res, t)
				tmp = tmp[:len(tmp)-1]
				continue
			}
			f(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	f(1, []int{})
	return res
}

func combineOfficial(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}

func combineOfficial2(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
