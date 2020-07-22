package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(INTMAX)
	x := 123
	x = 1563847412
	x = -1563847412
	x = 9463847412
	x = -9463847412
	fmt.Println(reverseOfficial(x))
}

const INTMAX = 1 << 31

func reverse(x int) int {
	xArr := strings.Split(strconv.Itoa(x), "")
	yArr := make([]string, len(xArr))
	var start, index int = 0, 1
	if xArr[0] == "-" {
		start = 1
		index = 0
		yArr[0] = "-"
	}

	for i := start; i < len(xArr); i++ {
		yArr[len(xArr)-index-i] = xArr[i]
	}

	yStr := strings.Join(yArr, "")
	y, _ := strconv.Atoi(yStr)
	if y < INTMAX-1 && y > -INTMAX {
		return y
	} else {
		return 0
	}
}

func reverseOfficial(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > INTMAX/10 || (rev == INTMAX/10 && pop > 7) {
			return 0
		}
		if rev < -INTMAX/10 || (rev == -INTMAX/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}

	return rev
}
