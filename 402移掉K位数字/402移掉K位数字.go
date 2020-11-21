package main

import (
	"fmt"
	"strings"
)

func main() {
	num := "1432219"
	//num = "23401"
	k := 3

	//num = "10010200"
	//k = 1

	//num = "10"
	//k = 1
	//num = "10200"
	//num = "1107"
	num = "112"
	//num = "121"
	//num = "214"
	//num = "210"
	k = 1

	//num = "4321"
	//k = 2
	fmt.Println(removeKdigits(num, k))
	//fmt.Println(removeKdigits0(num, k))
	//fmt.Println(removeKdigits1(num, k))
}

//8 ms-25.99%	3.1 MB-30.56%
func removeKdigits1(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	remain := n - k //需要保留的字符串长度

	nums := []byte(num)
	var arr []byte //记录需要保留的字符
	var dfs func(index, remain int)
	dfs = func(index, remain int) {
		if remain == 0 {
			return
		}
		var temp = byte('0' + 10)
		var position int
		var myRange = len(nums) - (remain - 1)
		for i := index; i < myRange; i++ {
			if nums[i] < temp {
				temp = nums[i]
				position = i
			}
		}
		arr = append(arr, temp)
		if remain == len(nums)-position {
			arr = append(arr, nums[position+1:]...)
			return
		}
		dfs(position+1, remain-1)
	}
	dfs(0, remain)
	ans := strings.TrimLeft(string(arr), "0")
	if len(ans) == 0 {
		return "0"
	}

	return ans
}

//8 ms-25.99%	2.6 MB-86.63%
func removeKdigits0(num string, k int) string {
	length := len(num)
	step := length - k
	var ans []byte
	minPos := -1
	for step > 0 {
		start := minPos + 1
		minVal := num[start]
		for i := length - step; i >= start; i-- {
			if num[i] <= minVal {
				minVal = num[i]
				minPos = i
			}
		}
		if minVal == '0' && len(ans) == 0 {
			step--
		} else {
			ans = append(ans, minVal)
			step--
		}
	}
	if len(ans) == 0 {
		return "0"
	}
	return string(ans)
}

//0 ms	2.6 MB
func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
