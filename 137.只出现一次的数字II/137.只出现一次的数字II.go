package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 2}
	nums = []int{0, 1, 0, 1, 0, 1, 99}

	fmt.Println(singleNumber(nums))
}

//8 ms-63.43%	3.4 MB-93.69%
func singleNumber(nums []int) int {
	var number, ans int
	for i := 0; i < 64; i++ {
		//初始化每一位1的个数为0
		number = 0
		for _, num := range nums {
			//通过右移i位的方式，计算每一位1的个数
			number += (num >> i) & 1
		}
		//最终将抵消后剩余的1放在对应的位数上
		ans |= (number) % 3 << i
	}
	return ans
}
