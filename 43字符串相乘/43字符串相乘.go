package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num1 := "923"
	num2 := "45"

	//num1 = "123456789"
	//num2 = "987654321"
	//
	//num1 = "2"
	//num2 = "3"
	//fmt.Println(multiply(num1, num2))
	//fmt.Println(multiplyOfficial(num1, num2))
	fmt.Println(multiplyOfficial2(num1, num2))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1 //交换位置，将大数放前面
	}

	m := len(num1)
	n := len(num2)
	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, m+n)
		for j := 0; j < m+n; j++ {
			matrix[i][j] = "0"
		}
	}

	for i := n - 1; i >= 0; i-- {
		var y int
		for j := m - 1; j >= 0; j-- {
			num := int(num2[i]-'0') * int(num1[j]-'0')
			x := (num + y) % 10 //余数，
			y = (num + y) / 10  //是否进一
			matrix[n-1-i][j+i+1] = strconv.Itoa(x)
			if j == 0 && y != 0 {
				matrix[n-1-i][j+i] = strconv.Itoa(y)
			}
		}
	}

	res := "0"
	for i := 0; i < n; i++ {
		res = addStringsOfficial(res, strings.Join(matrix[i], ""))
	}

	return strings.TrimLeft(res, "0")
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

func multiplyOfficial(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		curr := ""
		add := 0
		for j := n - 1; j > i; j-- {
			curr += "0"
		}
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			product := x*y + add
			curr = strconv.Itoa(product%10) + curr
			add = product / 10
		}
		for ; add != 0; add /= 10 {
			curr = strconv.Itoa(add%10) + curr
		}
		ans = addStringsOfficial(ans, curr)
	}
	return ans
}

func multiplyOfficial2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	ans := ""
	idx := 0
	if ansArr[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArr[idx])
	}
	return ans
}
