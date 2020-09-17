package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nums = []int{1, 2, 3, 4}
	//nums = []int{1}

	//fmt.Println(permute(nums))
	fmt.Println(permutedfs(nums))
	//fmt.Println("----------")
	//fmt.Println(permute000(nums))
}

//思想：一个数：1，二个数：12、21，三个数312、132、123、321、231、213
//可以看作是把后一次加入的数插入前一次生成的数组中
//0 ms-100.00%	2.8 MB-18.70%
func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n == 0 {
		return ans
	} else if n == 1 {
		return append(ans, nums)
	}
	ans = [][]int{
		{nums[0], nums[1]},
		{nums[1], nums[0]},
	}
	for k := 2; k < n; k++ {
		num := nums[k]
		ansLen := len(ans)
		for i := 0; i < ansLen; i++ {
			for j := 0; j <= len(ans[i]); j++ {
				tmp := []int{}
				if j == 0 {
					tmp = append(append(tmp, num), ans[i]...)
				} else {
					t := make([]int, len(ans[i][:j]))
					copy(t, ans[i][:j])
					tmp = append(append(t, num), ans[i][j:]...)
				}
				ans = append(ans, tmp)
			}
		}
		ans = ans[ansLen:]
	}

	return ans
}

//根据以前的代码写的，闹着玩的	4 ms	8.41%	2.8 MB	17.56%
func permute000(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n == 0 {
		return ans
	} else if n == 1 {
		return append(ans, nums)
	}

	factorialArr := factorial(n)
	t := factorialArr[n]
	for i := 0; i < t; i++ {
		arr := make([]int, n)
		copy(arr, nums)
		ans = append(ans, getPermutation(arr, factorialArr, n, i+1))
	}

	return ans
}

func getPermutation(nums, factorialArr []int, n, k int) []int {
	var arr []int
	var fun func(n, k int)

	fun = func(n, k int) {
		f := factorialArr[n-1]
		sh := k / f
		yu := k % f
		var index int
		if yu == 0 {
			index = sh - 1
			yu = k - index*f
		} else {
			index = sh
		}
		arr = append(arr, nums[index])
		nums = append(nums[:index], nums[index+1:]...)
		if len(nums) == 1 {
			arr = append(arr, nums[0])
			return
		}
		fun(n-1, yu)
	}

	fun(n, k)
	return arr
}

func factorial(n int) []int {
	factorialArr := []int{
		1,
	}
	facVal := 1
	for i := 1; i <= n; i++ {
		facVal *= i
		factorialArr = append(factorialArr, facVal)
	}
	return factorialArr
}

func permutedfs(nums []int) [][]int {
	n := len(nums)
	used := make([]bool, n)
	var ans [][]int
	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		if index == n {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			tmp = append(tmp, nums[i])
			dfs(index+1, tmp)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
	dfs(0, []int{})

	return ans
}
