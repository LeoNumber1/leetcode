package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//[[-1],[2,3],[1,-1,-3]]
	triangle := [][]int{
		//{2},
		//{3, 4},
		//{6, 5, 7},
		//{4, 1, 8, 3},
		//{4, 3, 8, 1, 5},
		{-1},
		{2, 3},
		{1, -1, -3},
	}
	fmt.Println(minimumTotal3(triangle))
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	var sum, index int
	for i := 0; i < len(triangle); i++ {
		if i == 0 {
			sum += triangle[0][0]
			continue
		}
		var min int
		if triangle[i][index] > triangle[i][index+1] {
			min = triangle[i][index+1]
			index++
		} else {
			min = triangle[i][index]
		}
		sum += min
	}
	return sum
}

func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	var sum, index int
	//len1 := 2^(len(triangle)-1)		总共2^(n-1)
	var res [][]int
	for i := 0; i < len(triangle); i++ {
		var res1 []int
		if i == 0 {
			sum += triangle[0][0]
			res1 = append(res1, sum)
			res = append(res, res1)
			continue
		}

		for j := 0; j < len(res[i-1]); j++ {
			//for z := j; z < len(triangle[i]); z++ {
			//
			//}
			index = j
			res1 = append(res1, res[i-1][j]+triangle[i][index])
			res1 = append(res1, res[i-1][j]+triangle[i][index+1])
		}
		res = append(res, res1)

		//			[4,1,8,3]
		//[[2] [5 6] [11 10 11 13] [15 12  11 18  12 19  21 16]]
		//							[4,3,8,1,5]

		//var min int = math.MaxInt32
		//for j := index; j < len(triangle[i]); j++ {
		//	if j > index+1 {
		//		break
		//	}
		//	if triangle[i][j] < min {
		//		min = triangle[i][j]
		//		index = j
		//	}
		//}
		//sum += min
	}
	fmt.Println(res)
	return sum
}

func minimumTotal3(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	var sum int = math.MaxInt32
	//len1 := 2^(len(triangle)-1)		总共2^(n-1)
	sep := "|"
	var res [][]string
	for i := 0; i < len(triangle); i++ {
		var index int
		var res1 []string
		if i == 0 {
			//sum += triangle[0][0]
			res1 = append(res1, fmt.Sprint(triangle[0][0], sep, index))
			res = append(res, res1)
			continue
		}

		for j := 0; j < len(res[i-1]); j++ {
			temp := strings.Split(res[i-1][j], sep)
			index, _ = strconv.Atoi(temp[1])
			num, _ := strconv.Atoi(temp[0])
			//res1 = append(res1, res[i-1][j]+triangle[i][index])
			res1 = append(res1, fmt.Sprint(num+triangle[i][index], sep, index))
			res1 = append(res1, fmt.Sprint(num+triangle[i][index+1], sep, index+1))
		}
		res = append(res, res1)

	}
	fmt.Println(res)
	for _, v := range res[len(res)-1] {
		temp := strings.Split(v, sep)
		num, _ := strconv.Atoi(temp[0])
		if num < sum {
			sum = num
		}
	}
	return sum
}

//			[4,1,8,3]
//[[2] [5 6] [11 10 11 13] [15 12  11 18  12 19  21 16]]
//							[4,3,8,1,5]
