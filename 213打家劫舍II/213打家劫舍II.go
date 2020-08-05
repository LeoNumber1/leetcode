package main

import (
	"fmt"
)

func main() {
	//[2,3,2]
	nums := []int{
		//2, 3, 2,
		//1, 2, 3, 1,
		//3, 2, 3, 4, 2, 3,
		2, 3,
	}

	fmt.Println(rob1(nums))
}

func rob(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	case 3:
		return max(max(nums[0], nums[1]), nums[2])
	}

	matrix := make([][]int, 3)
	//定义 3*n 的二维数组，第一行放置小偷偷到第几家时能偷取的最大值；第二行放置该最大值是否包含了第一个值，1为包含，0为不包含；第三行放置不包含第一个值的能偷取到的最大值
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, n)
	}

	matrix[0][0] = nums[0] //就一家，偷就对了
	matrix[1][0] = 1       //1为包含第一家
	matrix[2][0] = 0       //除了第一家，可偷取0

	if nums[0] > nums[1] { //第一家比第二家多
		matrix[0][1] = nums[0] //偷第一家
		matrix[1][1] = 1       //包含第一家
	} else { //第一家比第二家少
		matrix[0][1] = nums[1] //只偷第二家
		matrix[1][1] = 0       //不包含第一家
	}
	matrix[2][1] = nums[1] //不包含第一家，只能偷第二家

	if nums[0]+nums[2] > nums[1] {
		matrix[0][2] = nums[0] + nums[2]
		matrix[1][2] = 1
	} else {
		matrix[0][2] = nums[1]
		matrix[1][2] = 0
	}
	matrix[2][2] = max(nums[1], nums[2])

	for i := 3; i < n; i++ {
		if i == n-1 { //如果是最后一家，需要判断之前最大的有没有抢第一家
			var rubMax int
			if matrix[0][i-1] > matrix[0][i-2]+nums[i] {
				rubMax = matrix[0][i-1]
			} else {
				if matrix[1][i-2] == 0 { // 倒数第三家最大抢劫数不包括第一家
					rubMax = matrix[0][i-2] + nums[i]
				} else {
					rubMax = matrix[0][i-1]
				}
			}
			matrix[0][n-1] = max(rubMax, max(nums[i]+matrix[2][i-2], matrix[2][i-1]))
			break
		}
		//如果不是最后一家
		if matrix[0][i-1] > matrix[0][i-2]+nums[i] {
			matrix[0][i] = matrix[0][i-1]
			matrix[1][i] = matrix[1][i-1]
		} else {
			matrix[0][i] = matrix[0][i-2] + nums[i]
			matrix[1][i] = matrix[1][i-2]
		}
		matrix[2][i] = max(matrix[2][i-1], matrix[2][i-2]+nums[i])
	}
	return matrix[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[0] = nums[0]
	dp2[0] = 0
	dp1[1] = max(nums[0], nums[1])
	dp2[1] = nums[1]
	for i := 2; i < n; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return max(dp1[n-2], dp2[n-1])
}
