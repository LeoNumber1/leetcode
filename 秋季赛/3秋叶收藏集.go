package main

import (
	"fmt"
	"strings"
)

func main() {
	leaves := "rrryyyrryyyrr"
	//leaves = "rryyyrrrrryyyyyyyyrryyyrr"
	//leaves = "ryr"
	//leaves = "yyryy"
	//leaves = "rrryyyrrr"
	//leaves = "rrrrrrrrrr"
	//leaves = "yyyyy"

	//fmt.Println(minimumOperations(leaves))
	//fmt.Println(minimumOperations1(leaves))
	//fmt.Println(minimumOperations2(leaves))
	fmt.Println(minimumOperations3(leaves))
	fmt.Println(minimumOperations4(leaves))
	fmt.Println(minimumOperations5(leaves))
}

//这个可以，但超出时间限制了
func minimumOperations2(leaves string) int {
	var count int
	var r = []rune(leaves)
	if leaves[0] != 'r' {
		count++
		r[0] = 'r'
	}
	if leaves[len(r)-1] != 'r' {
		count++
		r[len(r)-1] = 'r'
	}

	leaves = strings.Trim(string(r), "r")
	//fmt.Println(leaves)
	n := len(leaves)
	if n == 0 { //中间没黄，翻一次
		return count + 1
	}

	arr := []int{} //奇数为r，偶数为y，值代表数量

	var pre byte
	var countR, countY int
	var totalR int
	for i := 0; i < n; i++ {
		if leaves[i] == 'r' {
			totalR++
		}
		if leaves[i] != pre {
			if leaves[i] == 'r' {
				arr = append(arr, 1)
				countR++
			} else {
				arr = append(arr, 1)
				countY++
			}
			pre = leaves[i]
		} else {
			arr[len(arr)-1]++
		}
	}

	if countR == 0 { //中间没红，不用翻面，返回0
		return count
	}

	var totalY = len(leaves) - totalR
	var count1 int = 100000 //

	for i := 0; i < len(arr); i++ {
		if i%2 == 0 { //偶数，y
			for j := i; j < len(arr); j = j + 2 {
				changeR, notChangeY := 0, 0
				for k := i; k <= j; k++ {
					if k%2 == 1 { //奇数，r
						changeR += arr[k]
					} else { //偶数，y
						notChangeY += arr[k]
					}
				}
				changeTotal := changeR + totalY - notChangeY

				count1 = min(count1, changeTotal)
			}
		}
	}

	return count + count1
}

func minimumOperations3(leaves string) int {
	var count int
	if leaves[0] != 'r' {
		count++
		leaves = "r" + leaves[1:]
	}
	if leaves[len(leaves)-1] != 'r' {
		count++
		leaves = leaves[:len(leaves)-1] + "r"
	}

	leaves = strings.Trim(leaves, "r")
	n := len(leaves)
	if n == 0 { //中间没黄，翻一次
		return count + 1
	}

	arr := []int{} //奇数为r，偶数为y，值代表数量

	var pre byte
	var countR int
	var totalR int
	for i := 0; i < n; i++ {
		if leaves[i] == 'r' {
			totalR++
		}
		if leaves[i] != pre {
			if leaves[i] == 'r' {
				arr = append(arr, 1)
				countR++
			} else {
				arr = append(arr, 1)
			}
			pre = leaves[i]
		} else {
			arr[len(arr)-1]++
		}
	}

	if countR == 0 { //中间没红，不用翻面，返回0
		return count
	}

	var totalY = len(leaves) - totalR
	var count1 int = 100000 //

	for i := 0; i < len(arr); i = i + 2 {
		changeR, notChangeY := 0, 0

		for j := i; j < len(arr); j = j + 2 {
			notChangeY += arr[j]
			if i != j {
				changeR += arr[j-1]
			}
			changeTotal := changeR + totalY - notChangeY

			count1 = min(count1, changeTotal)
		}
	}

	return count + count1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//动态规划
func minimumOperations4(leaves string) int {
	l := len(leaves)
	if l < 3 {
		return -1
	}
	var ans int
	//pre handle
	if leaves[0] == 'y' {
		ans++
	}
	if leaves[l-1] == 'y' {
		ans++
	}

	//开始处理中间数据
	//因为已经确保两边都有红了，所以中间随便怎么切两刀就好了
	//dp[][0]为全红，只能从全红变过来
	//dp[][1]红黄，能从红或者红黄变过来
	//dp[][2]红黄红，从红黄或者红黄红变过来
	var dp [][3]int
	for i := 0; i < l-1; i++ {
		dp = append(dp, [3]int{})
	}
	for i := 1; i < l-1; i++ {
		if leaves[i] == 'y' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
			if i > 1 {
				dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + 1
			} else {
				dp[i][2] = dp[i][1]
			}
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
			if i > 1 {
				dp[i][2] = min(dp[i-1][2], dp[i-1][1])
			} else {
				dp[i][2] = dp[i][1]
			}
		}
	}
	//不能全红
	return ans + min(dp[l-2][1], dp[l-2][2])
}

func minimumOperations5(leaves string) int {
	n := len(leaves)
	dp := [3][]int{}
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n)
	}
	// 动态规划
	// 状态定义： 字符串要满足ryr模式，有三个状态：
	// 1. 满足全部是r的最小调整次数
	// 2. 满足ry形式的最小调整次数
	// 3. 满足ryr形式的最小调整次数
	INF := int(1e8) // 定义无穷大
	if leaves[0] == 'r' {
		dp[0][0] = 0   //第一个字母是r,天然满足状态1，不需要操作
		dp[1][0] = INF //长度为1时不可能满足，花费无穷大
		dp[2][0] = INF //长度为1时不可能满足，花费无穷大
	} else {
		dp[0][0] = 1 //第一个字母是y
		dp[1][0] = INF
		dp[2][0] = INF
	}
	dp[2][1] = INF // 长度为2时不可能满足状态3

	for i := 1; i < n; i++ {
		if leaves[i] == 'r' {
			dp[0][i] = dp[0][i-1]                      //如果当前为r，则与上次一样
			dp[1][i] = min(dp[0][i-1], dp[1][i-1]) + 1 //满足状态2，考虑是从状态1转换或者上次的状态2转换哪个成本最小，再+1
			dp[2][i] = min(dp[1][i-1], dp[2][i-1])
		} else {
			dp[0][i] = dp[0][i-1] + 1 //需要操作一次，最后一位y变成r
			dp[1][i] = min(dp[0][i-1], dp[1][i-1])
			dp[2][i] = min(dp[1][i-1], dp[2][i-1]) + 1
		}
	}
	return dp[2][n-1]
}

//static public int minimumOperations(String leaves) {
//char[] chars = leaves.toCharArray();
//
//// 动态规划
//// 状态定义： 字符串要满足ryr模式，有三个状态：
//// 1. 满足全部是r的最小调整次数
//// 2. 满足ry形式的最小调整次数
//// 3. 满足ryr形式的最小调整次数
//// 三列 每列长度length
//int[][] dp = new int[3][chars.length];
//int INF = (int)1e8; // 定义无穷大
//
//if (chars[0] == 'r'){
//dp[0][0] = 0;   // 第一个字母是r,天然满足状态1，不需要操作
//dp[1][0] = INF;   // 长度为1时不可能满足，花费无穷大
//dp[2][0] = INF;
//}else {
//dp[0][0] = 1;
//dp[1][0] = INF;   // 第一个字母是r,调整到状态1，操作1次
//dp[2][0] = INF;
//}
//dp[2][1] = INF; // 长度为2时不可能满足状态3

//for (int i = 1; i < chars.length; i++) {
//if (chars[i] == 'r'){
//// 如下一个字母是r
//dp[0][i] = dp[0][i-1];  // 不需要操作，和上一次状态1的次数相同
//dp[1][i] = Math.min(dp[0][i-1],dp[1][i-1])+1; // 考虑要满足状态2， 是在上一次转化为状态1时花费少还是全转化为状态2时花费少 选最小花费加上本次操作数1
//dp[2][i] = Math.min(dp[1][i-1],dp[2][i-1]); // 考虑要满足状态3， 是在上一次状态为2时花费少还是状态为3时花费少 本次操作没有花费
//}else {
//dp[0][i] = dp[0][i-1]+1;  // 需要操作1次，即最后一位由y改成r
//dp[1][i] = Math.min(dp[0][i-1],dp[1][i-1]); // 考虑要满足状态2， 是在上一次转化为状态1时花费少还是全转化为状态2时花费少 本次没有花费
//dp[2][i] = Math.min(dp[1][i-1],dp[2][i-1])+1; // 考虑要满足状态3， 是在上一次状态为2时花费少还是状态为3时花费少 选最小花费加上本次操作数1
//}
//}
//
//// 返回最大长度下 状态3的最小花费数
//return dp[2][chars.length-1];
//}
