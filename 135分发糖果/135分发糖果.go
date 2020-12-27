package main

import "fmt"

func main() {
	ratings := []int{1, 0, 2} //5
	//ratings = []int{1, 2, 2}          //4
	//ratings = []int{2, 1, 0, 2}       //8
	//ratings = []int{3, 2, 2, 1, 0}    //9
	//ratings = []int{1, 3, 2, 2, 1}    //7
	ratings = []int{1, 2, 3, 2, 1, 0} //13
	//ratings = []int{1, 2, 3, 2, 1}    //9
	//ratings = []int{1, 2, 3, 1, 0}    //9
	ratings = []int{1, 3, 2, 1, 0} //11
	//ratings = []int{1, 3, 2, 0} //7
	//ratings = []int{1, 3, 4, 5, 2}    //11

	//fmt.Println(candy(ratings))
	fmt.Println(candyOfficial1(ratings))
	fmt.Println(candyOfficial2(ratings))
}

//28 ms-34.32%	6.4 MB-39.26%   我自己写的动态规划
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n) //dp[i][j]，i为数字下标，j=0为当前小朋友糖数，j=1为总糖数
	dp[0] = [2]int{1, 1}
	for i := 1; i < n; {
		if ratings[i] > ratings[i-1] {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = dp[i-1][1] + dp[i][0]
			i++
		} else if ratings[i] < ratings[i-1] {
			final := 0
			ii := i - 1
			for j := i; j < n && ratings[j] < ratings[j-1]; j++ {
				i++
				final = j
			}
			temp := 1
			for j := final; j >= ii; j-- {
				if dp[j][0] > temp {
					break
				}
				dp[j][0] = temp
				temp++
			}
			for j := ii; j <= final; j++ {
				if j == 0 {
					dp[j][1] = dp[j][0]
					continue
				}
				dp[j][1] = dp[j-1][1] + dp[j][0]
			}
		} else {
			dp[i][0] = 1
			dp[i][1] = dp[i-1][1] + 1
			i++
		}
	}
	return dp[n-1][1]
}

//16 ms-96.95%	6.2 MB-51.57%   官方的两次遍历
func candyOfficial1(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

//20 ms-81.44%	6.1 MB-100.00%  官方的一次遍历
func candyOfficial2(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
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
