package main

import "fmt"

func main() {
	ratings := []int{1, 0, 2}         //5
	ratings = []int{1, 2, 2}          //4
	ratings = []int{2, 1, 0, 2}       //8
	ratings = []int{3, 2, 2, 1, 0}    //9
	ratings = []int{1, 3, 2, 2, 1}    //7
	ratings = []int{1, 2, 3, 2, 1, 0} //13
	//ratings = []int{1, 2, 3, 2, 1}    //9
	//ratings = []int{1, 2, 3, 1, 0}    //9
	//ratings = []int{1, 3, 2, 1, 0}    //11
	//ratings = []int{1, 3, 2, 0}       //7
	//ratings = []int{1, 3, 4, 5, 2}    //11

	fmt.Println(candy(ratings))
}

//有问题
func candy0(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n) //dp[i][j]，i为数字下标，j=0为当前小朋友糖数，j=1为总糖数
	dp[0] = [2]int{1, 1}
	index := -1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = dp[i-1][1] + dp[i][0]
			index = -1
		} else if ratings[i] < ratings[i-1] {
			if index == -1 {
				index = i - 1
			}
			if ratings[i] < ratings[i-1]-1 {
				index = i
			}
			temp := 0
			dp[i][0] = 1
			if dp[i-1][0] == 1 {
				temp = i - index
			}
			dp[i][1] = dp[i-1][1] + dp[i][0] + temp
		} else {
			dp[i][0] = 1
			dp[i][1] = dp[i-1][1] + 1
			index = -1
		}
	}
	return dp[n-1][1]
}

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
