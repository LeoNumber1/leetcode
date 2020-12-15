package main

import (
	"fmt"
	"strconv"
)

func main() {
	N := 1234
	N = 1231
	N = 110
	N = 9

	//fmt.Println(monotoneIncreasingDigits(N))
	fmt.Println(monotoneIncreasingDigitsOfficial(N))
}

//0 ms-100.00%	2 MB-65.71%
func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	num := []byte(strconv.Itoa(N))
	index, is := isIncrease(num)
	if is {
		return N
	}
	num[index]--
	for i := index + 1; i < len(num); i++ {
		num[i] = '9'
	}

	n, _ := strconv.Atoi(string(num))
	return n
}

func isIncrease(num []byte) (int, bool) {
	index, is := 0, true
	for i := 1; i < len(num); i++ {
		if num[i] < num[index] {
			is = false
			break
		} else if num[i] > num[index] {
			index = i
		}
	}
	return index, is
}

//0 ms-100.00%	2 MB-80.00%
func monotoneIncreasingDigitsOfficial(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
