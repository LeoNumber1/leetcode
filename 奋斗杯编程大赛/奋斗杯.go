package main

import "fmt"

func main() {
	//详细题目说明可见：https://blog.csdn.net/u012140251/article/details/109409015
	//-2进制表示法
	//n := -10
	//intToNegativeBinary(n)
	//intToNegativeBinary(1)
	//for i := -10; i < 10; i++ {
	//	intToNegativeBinary(i)
	//}

	//详细题目说明可见：https://blog.csdn.net/u012140251/article/details/109410044
	//跑步，走路不得分，慢跑得一倍分，快跑得两倍分，但下一段只能走路，求最大得分
	//如4段，每段分值 1 2 3 4 ， 则最大得分为 1+2+3+4*2 = 14
	fmt.Println(run([]int{1, 2, 3, 4}))
	fmt.Println(runBetter([]int{1, 2, 3, 4}))
	//fmt.Println(run([]int{4, 3, 2, 1}))
	//fmt.Println(runBetter([]int{4, 3, 2, 1}))
	//fmt.Println(run([]int{4, 3, 1, 2}))
	//fmt.Println(runBetter([]int{4, 3, 1, 2}))
}

func intToNegativeBinary(n int) {
	if n == 0 {
		fmt.Println(0)
		return
	}
	const BASE = -2
	var ans []byte
	for n != 1 {
		if n > 0 {
			ans = append(ans, byte('0'+n%BASE))
			n /= BASE
		} else {
			yu := n % BASE
			if yu == -1 {
				//余数为-1
				n = n/BASE + 1
				ans = append(ans, '0'+byte(1))
			} else {
				ans = append(ans, '0'+byte(yu))
				n /= BASE
			}
		}
	}
	//把最后剩下的1加进去
	ans = append(ans, '0'+byte(1))

	length := len(ans)
	//反转
	for i := 0; i < length/2; i++ {
		ans[i], ans[length-1-i] = ans[length-1-i], ans[i]
	}
	fmt.Println(string(ans))
}

func run(points []int) int {
	length := len(points)
	if length == 0 {
		return 0
	}

	//dp[i][j]值为当前路段当前状态获得的最大总分数，i为路段的索引，j为本段状态，其中0为本段走路，1为本段慢跑，2为本段快跑
	dp := make([][3]int, length)
	dp[0][0] = 0
	dp[0][1] = points[0]
	dp[0][2] = points[0] * 2
	for i := 1; i < length; i++ {
		dp[i][0] = max(max(dp[i-1][0], dp[i-1][1]), dp[i-1][2]) //本段走路，上段可以是任何状态
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + points[i]      //本段慢跑，上段只能是走路或者慢跑
		dp[i][2] = max(dp[i-1][0], dp[i-1][1]) + points[i]*2    //本段快跑，上段只能是走路或者慢跑
	}
	return max(max(dp[length-1][0], dp[length-1][1]), dp[length-1][2])
}

func runBetter(points []int) int {
	length := len(points)
	if length == 0 {
		return 0
	}

	var prevWalk, prevSlowRun, prevFastRun = 0, points[0], points[0] * 2
	for i := 1; i < length; i++ {
		var walk, slowRun, fastRun int
		walk = max(max(prevWalk, prevSlowRun), prevFastRun) //本段走路，上段可以是任何状态
		slowRun = max(prevWalk, prevSlowRun) + points[i]    //本段慢跑，上段只能是走路或者慢跑
		fastRun = max(prevWalk, prevSlowRun) + points[i]*2  //本段快跑，上段只能是走路或者慢跑
		prevWalk = walk
		prevSlowRun = slowRun
		prevFastRun = fastRun
	}
	return max(max(prevWalk, prevSlowRun), prevFastRun)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
