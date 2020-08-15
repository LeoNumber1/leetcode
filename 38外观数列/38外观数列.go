package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	n := 43
	//fmt.Println(countAndSay(n))
	//fmt.Println(countAndSay1(n))
	t2 := time.Now()
	countAndSay2(n)
	fmt.Println(time.Since(t2))
	t3 := time.Now()
	countAndSay3(n)
	fmt.Println(time.Since(t3))
	t4 := time.Now()
	countAndSay4(n)
	fmt.Println(time.Since(t4))
}

func countAndSay(n int) string {
	f := make([]string, n)
	f[0] = "1"
	for i := 1; i < n; i++ {
		for j := 0; j < len(f[i-1]); j++ {
			count := 1
			for k := j + 1; k < len(f[i-1]); k++ {
				if f[i-1][j] == f[i-1][k] {
					count++
					j++
				} else {
					break
				}
			}
			f[i] += strconv.Itoa(count) + string(f[i-1][j])
		}
	}
	return f[n-1]
}

func countAndSay1(n int) string {
	str := "1"
	var res string
	for i := 1; i < n; i++ {
		for j := 0; j < len(str); j++ {
			count := 1
			for k := j + 1; k < len(str); k++ {
				if str[j] == str[k] {
					count++
				} else {
					break
				}
			}
			j += count - 1
			res += strconv.Itoa(count) + string(str[j])
		}
		str = res
		res = ""
	}
	return str
}
func countAndSay2(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		var res string
		prebyte := str[0]
		count := 1
		for j := 1; j < len(str); j++ {
			if str[j] == prebyte {
				count++
			} else {
				res += strconv.Itoa(count) + string(prebyte)
				count = 1
				prebyte = str[j]
			}
		}
		res += strconv.Itoa(count) + string(prebyte)
		str = res
	}
	return str
}

func countAndSay3(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		var tmp strings.Builder
		preByte := str[0]
		count := 1
		for j := 1; j < len(str); j++ {
			if str[j] == preByte {
				count++
			} else {
				tmp.WriteString(strconv.Itoa(count))
				tmp.WriteByte(preByte)
				preByte = str[j]
				count = 1
			}
		}
		tmp.WriteString(strconv.Itoa(count))
		tmp.WriteByte(preByte)
		str = tmp.String()
	}
	return str
}

func countAndSay4(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		var tmp string
		preByte := str[0]
		count := 1
		for j := 1; j < len(str); j++ {
			if str[j] == preByte {
				count++
			} else {
				tmp += strconv.Itoa(count) + string(preByte)
				preByte = str[j]
				count = 1
			}
		}
		tmp += strconv.Itoa(count) + string(preByte)
		str = tmp
	}
	return str
}
