package main

import "fmt"

func main() {
	N := 6
	fmt.Println(divisorGame2(N))
}

func divisorGame(N int) bool {
	return N%2 == 0
}

func divisorGame2(N int) bool {
	f := make([]bool, N+1) //定义 f[i] 表示当前数字 i 的时候先手是处于必胜态还是必败态
	f[1] = false
	f[2] = true

	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ { //j是Alice取的值
			if i%j == 0 && !f[i-j] {
				f[i] = true
				break
			}
		}
	}
	return f[N]
}
