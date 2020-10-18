package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5}
	//nums = []int{1, 5, 11, 5, 4}
	//nums = []int{11, 1}

	fmt.Println(canPartition(nums))
}

//dp 动态规划、空间优化  8 ms	2.5 MB
func canPartition1(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 != 0 {
		return false
	}

	capacity := sum / 2
	var dp = make([]bool, capacity+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := capacity; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
		if dp[capacity] {
			return true
		}
	}
	return dp[capacity]
}

//dp 动态规划   20 ms	6.6 MB
func canPartition0(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 != 0 {
		return false
	}

	capacity := sum / 2
	//创建二维状态数组，行：物品索引，列：容量（包括0）
	dp := make([][]bool, len(nums))

	for k := range dp {
		dp[k] = make([]bool, capacity+1)
	}

	dp[0][0] = true

	//先填表格第0行，第一个数只能让容积为它自己的背包恰好装满
	if nums[0] <= capacity {
		dp[0][nums[0]] = true
	}
	//再填表格后面几行
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= capacity; j++ {
			//直接从上一行把结果先抄下来，然后再修正
			dp[i][j] = dp[i-1][j]
			if dp[i][j] {
				continue
			}
			if nums[i] == j {
				dp[i][j] = true
			} else if nums[i] < j {
				//dp[i][j]的值：不选nums[i]就是dp[i-1][j]，选nums[i]就是dp[i-1][j-nums[i]]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
			if dp[i][capacity] {
				return true
			}
		}
	}
	return dp[len(nums)-1][capacity]
}

//位运算	4 ms-95.00%	2.5 MB-56.31%
func canPartition(nums []int) bool {
	var sum, maxNum int
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	part := sum >> 1
	if sum&1 != 0 || len(nums) < 2 || maxNum > part {
		return false
	}
	var bits = make([]byte, part+1)
	bits[0] = 1
	for i := 0; i < len(nums); i++ {
		//size := part - nums[i]
		//for j := size; j >= 0; j-- {
		//	bits[j+nums[i]] |= bits[j]
		//}
		for j := part; j >= nums[i]; j-- {
			bits[j] |= bits[j-nums[i]]
		}
		//判断中位数如果是1，直接返回true
		if bits[part] == 1 {
			return true
		}
	}

	return false
}
