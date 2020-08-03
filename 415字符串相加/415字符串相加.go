package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "123128"
	num2 := "12345"

	num1 = "10000008889"
	num2 = "1111"

	//num1 = "8889123123213123123"
	//num2 = "8889123123213123123"
	//fmt.Println(num1[0])
	//fmt.Println(addStrings2(num1, num2))
	fmt.Println(addStringsOfficial(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	num1Len := len(num1)
	num2Len := len(num2)
	m := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	var x, y int
	var res string

	for i, j := num1Len-1, num2Len-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		num := m[string(num1[i])] + m[string(num2[j])] + x
		y = num % 10 //余数
		x = num / 10 //是否进1

		res = strconv.Itoa(y) + res
	}

	var len int
	var nums string
	if num1Len > num2Len {
		len = num1Len - num2Len
		nums = num1
	} else {
		len = num2Len - num1Len
		nums = num2
	}

	for i := len - 1; i >= 0; i-- {
		sum := m[string(nums[i])] + x
		y = sum % 10 //余数
		x = sum / 10 //是否进1

		res = strconv.Itoa(y) + res
	}

	if x != 0 {
		res = "1" + res
		return res
	}

	return res
}

func addStrings2(num1 string, num2 string) string {
	num1Len := len(num1)
	num2Len := len(num2)

	var x, y int
	var res string

	for len(num1) > 0 && len(num2) > 0 {
		num := int(num1[len(num1)-1]-'0') + int(num2[len(num2)-1]-'0') + x
		y = num % 10 //余数
		x = num / 10 //是否进1

		res = strconv.Itoa(y) + res
		num1 = num1[:len(num1)-1]
		num2 = num2[:len(num2)-1]
	}

	var len int
	var nums string
	if num1Len > num2Len {
		len = num1Len - num2Len
		nums = num1
	} else {
		len = num2Len - num1Len
		nums = num2
	}

	if x == 0 {
		res = nums + res
		return res
	}

	for i := len - 1; i >= 0; i-- {
		sum := int(nums[i]-'0') + x
		y = sum % 10 //余数
		x = sum / 10 //是否进1

		res = strconv.Itoa(y) + res

		if x == 0 {
			res = nums[:i] + res
			break
		}
	}

	if x != 0 {
		res = "1" + res
		return res
	}

	return res
}

func addStringsOfficial(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}
