package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	ring := "godding"
	key := "gd"

	//ring = "godaacceabding"
	//key = "gdb"
	//
	//ring = "pqwcx"
	//key = "cpqwx"
	//
	//ring = "xrrakuulnczywjs"
	//key = "jrlucwzakzussrlckyjjsuwkuarnaluxnyzcnrxxwruyr"
	//
	ring = "caotmcaataijjxi" //没排序优化前35.75s  排序优化后30.18s
	key = "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"
	//
	//ring = "caotmcaataijjxi"
	//key = "oatjii"

	//t := time.Now()
	//fmt.Println(findRotateSteps(ring, key), time.Since(t))
	t1 := time.Now()
	fmt.Println(findRotateStepsOfficial1(ring, key), time.Since(t1))
}

//时间超了
func findRotateSteps(ring string, key string) int {
	n := len(ring)
	arr := make([]map[byte][][2]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make(map[byte][][2]int) //定义二维数组，第一维是索引，第二维是距离
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := min1(j-i, n-j+i)
			arr[i][ring[j]] = append(arr[i][ring[j]], [2]int{j, temp})
			if i != j {
				arr[j][ring[i]] = append(arr[j][ring[i]], [2]int{i, temp})
			}
		}
	}

	for _, v := range arr {
		for _, v1 := range v {
			sort.Slice(v1, func(i, j int) bool {
				return v1[i][1] < v1[j][1]
			})
		}
	}

	keyLen := len(key)
	var rotate int = math.MaxInt32
	var dfs func(index, p, r int)
	dfs = func(index, p, r int) {
		if r >= rotate {
			return
		}
		if p == keyLen {
			if r < rotate {
				rotate = r
			}
			return
		}

		if p < keyLen-1 {
			for _, v := range arr[index][key[p+1]] {
				dfs(v[0], p+1, r+v[1])
			}
		} else {
			dfs(0, p+1, r)
		}
	}
	dfs(0, -1, 0)

	return rotate + keyLen
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//8 ms	6.6 MB
func findRotateStepsOfficial(ring string, key string) int {
	const inf = math.MaxInt64 / 2
	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([][]int, m) //定义 dp[i][j] 表示从前往后拼写出 key 的第 i 个字符，ring 的第 j 个字符与 12:00 方向对齐的最少步数（下标均从 0 开始）。
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
				//下一个位置只能从上一个字母位置转过来，找它的最小转到步数。+1是按一次
			}
		}
	}
	return min(dp[m-1]...)
}

//8 ms	6.5 MB
func findRotateStepsOfficial1(ring string, key string) int {
	const inf = math.MaxInt64 / 2
	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([]int, n) //定义 dp[i][j] 表示从前往后拼写出 key 的第 i 个字符，ring 的第 j 个字符与 12:00 方向对齐的最少步数（下标均从 0 开始）。
	dp0 := make([]int, n)
	for i := range dp {
		dp[i] = inf
		dp0[i] = inf
	}
	for _, p := range pos[key[0]-'a'] {
		dp[p] = min(p, n-p) + 1
	}
	for i := 1; i < m; i++ {
		temp := make([]int, n)
		copy(temp, dp)
		copy(dp, dp0)
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[j] = min(dp[j], temp[k]+min(abs(j-k), n-abs(j-k))+1)
				//下一个位置只能从上一个字母位置转过来，找它的最小转到步数。+1是按一次
			}
		}
	}
	return min(dp...)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
