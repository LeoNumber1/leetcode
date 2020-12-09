package main

import (
	"fmt"
)

func main() {
	n := 5

	fmt.Println(concatenatedBinary(n))
}

const MOD = 1000000007

func concatenatedBinary(n int) int {
	var ans int = 1
	right := 4
	length := 2
	for i := 2; i < n+1; i++ {
		if i >= right {
			right <<= 1
			length++
		}
		ans *= 1 << length
		ans += i
		ans %= MOD
	}
	return ans
}

func concatenatedBinary1(n int) (ans int) {
	for i := 1; i <= n; i++ {
		//ans = (ans<<bits.Len(uint(i)) | i) % (1e9 + 7)
	}
	return
}
