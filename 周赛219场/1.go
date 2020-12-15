package main

import "fmt"

func main() {
	n := 14
	fmt.Println(numberOfMatches(n))
}

func numberOfMatches(n int) int {
	var ans int
	for n > 1 {
		if n&1 == 0 {
			n /= 2
			ans += n
		} else {
			n /= 2
			ans += n + 1
		}
	}
	return ans
}
