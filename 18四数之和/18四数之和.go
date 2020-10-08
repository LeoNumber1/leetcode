package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2, 0}
	target := 0

	fmt.Println(fourSumOfficial1(nums, target))
}

//876 ms-5.02%	2.9 MB-32.94%
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var temp []int
	var f func(int, int)
	f = func(index int, target int) {
		if len(temp) == 4 {
			if target == 0 {
				ans = append(ans, append([]int(nil), temp...))
			}
			return
		}
		if index == len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			f(i+1, target-nums[i])
			temp = temp[:len(temp)-1]
		}
	}
	f(0, target)
	return ans
}

func fourSumOfficial1(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
