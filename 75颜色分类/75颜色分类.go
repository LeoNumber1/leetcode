package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0, 1}
	nums = []int{1, 0, 2, 2, 0, 2, 1, 1, 0, 2, 0, 2, 1, 1, 0, 2, 0, 2, 1, 1}

	//sortColors(nums)
	sortColorsOfficial2(nums)
	fmt.Println(nums)
}

func sortColors1(nums []int) {
	var index0, index2 = 0, len(nums) - 1
	for i, j := 0, len(nums)-1; i < len(nums) && j > 0 && i < j; {
		if nums[j] == 2 {
			nums[j], nums[index2] = nums[index2], nums[j]
			j--
			index2 = j
			continue
		}
		if nums[j] == 0 {
			nums[j], nums[index0] = nums[index0], nums[j]
			index0++
			i = index0
			continue
		}
		if nums[i] == 0 {
			nums[i], nums[index0] = nums[index0], nums[i]
			i++
			if index0 == 0 {
				//	index0 = i
				//} else {
				index0++
			}
			continue
		}
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
			if nums[j] == 2 {
				j--
				index2 = j
			}
			if nums[i] == 0 {
				i++
				index0 = i
			}
		} else {
			i++
		}
	}
}

//0 ms	2 MB
func sortColors(nums []int) {
	var m = map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	var start int
	for i := 0; i < m[0]; i++ {
		nums[i] = 0
	}
	start += m[0]
	for i := 0; i < m[1]; i++ {
		nums[start+i] = 1
	}
	start += m[1]
	for i := 0; i < m[2]; i++ {
		nums[start+i] = 2
	}
}

func sortColorsOfficial(nums []int) {
	count0 := swapColors(nums, 0) //把0排到前面
	swapColors(nums[count0:], 1)  //nums[:count0]全都是0了，对剩下的1排前面
}

func swapColors(colors []int, target int) (count int) {
	for i, color := range colors {
		if color == target {
			colors[i], colors[count] = colors[count], colors[i]
			count++
		}
	}
	return
}

func sortColorsOfficial2(nums []int) {

}
