package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}

	//nums = []int{1, 3, 12, 0, 0, 0, 0, 0}

	//moveZeroes(nums)
	moveZeroesOfficial(nums)
	fmt.Println(nums)
}

//执行耗时:4 ms,击败了94.38% 的Go用户
//内存消耗:3.7 MB,击败了99.94% 的Go用户
func moveZeroes(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	var left, right = 0, 0
	for left < n {
		if nums[left] == 0 {
			if right == 0 {
				right = left + 1
			}
			for ; right < n; right++ {
				if nums[right] != 0 {
					nums[left], nums[right] = nums[right], nums[left]
					right++
					break
				} else if right == n-1 {
					return
				}
			}
		}
		left++
	}
}

//执行耗时:4 ms,击败了94.38% 的Go用户
//内存消耗:3.7 MB,击败了99.94% 的Go用户
func moveZeroesOfficial(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
