package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{
		4, 6, 7, 7,
	}

	//fmt.Println(findSubsequences(nums))
	fmt.Println(findSubsequencesOfficial2(nums))
}

func findSubsequences(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, nums[i])
		for j := i + 1; j < len(nums); j++ {
			if nums[j] >= nums[i] {
				if !find(ret, []int{nums[i], nums[j]}) {
					ret = append(ret, []int{nums[i], nums[j]})
					//tmp = append(tmp, []int{nums[i], nums[j]})
				}
				tmp = append(tmp, nums[j])
				//} else {
				//	if !find(ret, []int{nums[i], nums[j]}) {
				//		ret = append(ret, []int{nums[j], nums[i]})
				//		tmp = append(nums[j], tmp)
				//		//tmp = append(tmp, []int{nums[j], nums[i]})
				//	}
			}
			if len(tmp) > 1 && !find(ret, tmp) {
				ret = append(ret, tmp)
			}
		}
	}

	return ret
}

func find(ret [][]int, target []int) bool {
	for _, v := range ret {
		if len(v) != len(target) {
			continue
		}
		for kt, vt := range target {
			if v[kt] != vt {
				break
			} else {
				if kt == len(target)-1 {
					return true
				}
			}
		}
	}
	return false
}

var (
	n    int
	temp []int
)

func findSubsequencesOfficial(nums []int) [][]int {
	n = len(nums)
	ans := [][]int{}
	set := map[int]bool{}
	for i := 0; i < 1<<n; i++ {
		findSubsequences1(i, nums)
		hashValue := getHash(263, int(1e9+7))
		if check() && !set[hashValue] {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
			set[hashValue] = true
		}
	}
	return ans
}

func findSubsequences1(mask int, nums []int) {
	temp = []int{}
	for i := 0; i < n; i++ {
		if (mask & 1) != 0 {
			temp = append(temp, nums[i])
		}
		mask >>= 1
	}
}

func getHash(base, mod int) int {
	hashValue := 0
	for _, x := range temp {
		hashValue = hashValue*base%mod + (x + 101)
		hashValue %= mod
	}
	return hashValue
}

func check() bool {
	for i := 1; i < len(temp); i++ {
		if temp[i] < temp[i-1] {
			return false
		}
	}
	return len(temp) >= 2
}

var (
	//temp []int
	ans [][]int
)

func findSubsequencesOfficial2(nums []int) [][]int {
	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

func dfs(cur, last int, nums []int) {
	if cur == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}
	if nums[cur] >= last {
		temp = append(temp, nums[cur])
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	if nums[cur] != last {
		dfs(cur+1, last, nums)
	}
}
