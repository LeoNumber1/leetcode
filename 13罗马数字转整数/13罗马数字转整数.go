package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "III"
	s = "IV"
	s = "IX"
	//s = "LVIII"
	s = "MCMXCIV"
	fmt.Println(romanToInt(s))
	fmt.Println(romanToInt1(s))
	fmt.Println(romanToInt2(s))
}

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	var num int
	sArr := strings.Split(s, "")
	n := len(sArr)
	for k := 0; k < n; k++ {
		if i, ok := m[sArr[k]]; ok {
			if k+1 < n && (i == 1 || i == 10 || i == 100) {
				if j, ok := m[sArr[k]+sArr[k+1]]; ok {
					num += j
					k = k + 1
				} else {
					num += i
				}
			} else {
				num += i
			}
		} else {
			return 0
		}
	}
	return num
}

func romanToInt1(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pre := 0
	var num int
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur >= pre {
			num += cur
		} else {
			num -= cur
		}
		pre = cur
	}
	return num
}

func romanToInt2(s string) int {
	pre := 0
	var num int
	for i := len(s) - 1; i >= 0; i-- {
		cur := getInt(s[i])
		if cur >= pre {
			num += cur
		} else {
			num -= cur
		}
		pre = cur
	}
	return num
}

func getInt(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
