package main

import "fmt"

func main() {
	stones := []int{5, 3, 1, 4, 2}
	//stones = []int{7, 90, 5, 1, 100, 10, 10, 2}

	fmt.Println(stoneGameVII(stones))
}

//有问题
func stoneGameVII0(stones []int) int {
	n := len(stones)
	left, right := 0, n-1
	if n == 2 {
		return max(stones[left], stones[right])
	}
	var total int
	for _, v := range stones {
		total += v
	}
	var ali, bob int
	var aliNow bool = true
	for total > 0 {
		if left == right {
			break
		}
		if stones[left] > stones[right] {
			if aliNow {
				total -= stones[right]
				right--
				ali += total
				aliNow = false
			} else {
				if right-left > 1 {
					total -= stones[left]
					left++
				} else {
					total -= stones[right]
					right--
				}
				bob += total
				aliNow = true
			}
		} else if stones[left] < stones[right] {
			if aliNow {
				total -= stones[left]
				left++
				ali += total
				aliNow = false
			} else {
				if right-left > 1 {
					total -= stones[right]
					right--
				} else {
					total -= stones[left]
					left++
				}
				bob += total
				aliNow = true
			}
		}
	}

	return ali - bob
}

func stoneGameVII1(stones []int) int {
	n := len(stones)
	pre := make([]int, n+1)
	dp := make([][]int, n)
	for k, stone := range stones {
		dp[k] = make([]int, n)
		pre[k+1] = pre[k] + stone
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = 0
				continue
			}
			L := pre[j+1] - pre[i+1] - dp[i+1][j]
			R := pre[j] - pre[i] - dp[i][j-1]
			dp[i][j] = max(L, R)
		}
	}
	return dp[0][n-1]
}

func stoneGameVII(stones []int) int {
	n := len(stones)
	dp := make([][]int, n)  //dp[i][j] : 表示从 i 到 j 的区间和
	res := make([][]int, n) //res[i][j]:表示轮到这一个人选时，石头剩下 i 到 j ，他能获得的最大价值差
	for k, stone := range stones {
		dp[k] = make([]int, n)
		res[k] = make([]int, n)
		for j := k; j < n; j++ {
			if k == j {
				dp[k][j] = stone
			} else {
				dp[k][j] = dp[k][j-1] + stones[j]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j-i == 1 {
				res[i][j] = max(stones[i], stones[j])
			} else {
				res[i][j] = max(dp[i+1][j]-res[i+1][j], dp[i][j-1]-res[i][j-1])
			}
		}
	}

	return res[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
