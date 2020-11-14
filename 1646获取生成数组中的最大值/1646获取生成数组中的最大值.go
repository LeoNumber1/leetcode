package main

import "fmt"

func main() {
	n := 7
	//n = 3
	fmt.Println(getMaximumGenerated(n))
}

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	var ans int = 1
	for i := 1; i <= n/2; i++ {
		a := 2 * i
		if a >= 2 && a <= n {
			nums[2*i] = nums[i]
			ans = max(nums[2*i], ans)
		}
		if a+1 >= 2 && a+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
			ans = max(nums[2*i+1], ans)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
