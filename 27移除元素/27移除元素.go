package main

import "fmt"

func main() {
	nums := []int{
		//3, 2, 2, 3,
		0, 1, 2, 2, 3, 0, 4, 2,
	}

	//fmt.Println(removeElement(nums, 2))
	fmt.Println(removeElement2(nums, 2))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

//双指针解法
func removeElement2(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 0
	j := 0
	for j < length {
		if nums[j] == val {
			// 去找一个不是 val 的值
			j++
		} else {
			// 赋值
			nums[i] = nums[j]
			i++
			j++
		}
	}

	return i
}
