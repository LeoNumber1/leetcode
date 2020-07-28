package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 10
	//x = 121
	x = 1231891321
	fmt.Println(isPalindrome(x))
	fmt.Println(isPalindromeString(x))
	fmt.Println(isPalindromeOfficial1(x))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var res int
	xBefore := x

	for x > 0 {
		y := x % 10
		x /= 10
		res *= 10
		res += y
	}
	return res == xBefore
}

func isPalindromeString(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	xStr := strconv.Itoa(x)
	n := len(xStr)
	for i := 0; i < n/2; i++ {
		if xStr[i] != xStr[n-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeOfficial1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNumber int
	for revertedNumber < x {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return revertedNumber == x || revertedNumber/10 == x
}
