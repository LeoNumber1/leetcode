package main

import "fmt"

func main() {
	x := 10
	//x = 121
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
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
