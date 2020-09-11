package main

import "fmt"

func main() {
	k := 3
	n := 7

	//n = 9
	//
	//k = 2
	//n = 12

	//fmt.Println(combinationSum3(k, n))
	fmt.Println(combinationSum3Official1(k, n))
}

func combinationSum3(k int, n int) (ans [][]int) {
	if n < k || n > 45 || n < 1 {
		return
	}

	var comb []int
	var dfs func(start, target int)
	dfs = func(start, target int) {
		if len(comb) > k || (len(comb) == k && target != 0) {
			return
		}
		if target == 0 && len(comb) == k {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		for i := start; i < 10; i++ {
			if target-i >= 0 {
				if len(comb) == k-1 && target > 9 {
					break
				}
				comb = append(comb, i)
				dfs(i+1, target-i)
				comb = comb[:len(comb)-1]
			} else {
				return
			}
		}
	}

	dfs(1, n)
	return
}

func combinationSum3Official1(k int, n int) (ans [][]int) {
	var temp []int
	check := func(mask int) bool {
		temp = nil
		sum := 0
		for i := 0; i < 9; i++ {
			if 1<<i&mask > 0 {
				temp = append(temp, i+1)
				sum += i + 1
			}
		}
		return len(temp) == k && sum == n
	}

	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			ans = append(ans, append([]int(nil), temp...))
		}
	}
	return
}
