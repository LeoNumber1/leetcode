package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "LEETCODEISHIRINGS"
	numRows := 4

	//s = "PAYPALISHIRING"
	//numRows = 3

	//fmt.Println(convert(s, numRows))
	fmt.Println(convertOfficial2(s, numRows))
}

//8 ms-79.23%	6.8 MB-14.54%
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	lenS := len(s)
	r := lenS / (numRows + numRows - 2) //循环次数
	n := r * (1 + numRows - 2)          //循环的列数
	var col int
	c := lenS % (numRows + numRows - 2)
	switch {
	case c == 0:
		col = n
	case c <= numRows:
		col = n + 1
	default:
		col = n + 1 + (lenS%(numRows+numRows-2) - numRows)
	}
	arr := make([][]byte, numRows)
	for k, _ := range arr {
		arr[k] = make([]byte, col)
	}

one:
	for j := 0; j < col; j++ {
		yu := j % (numRows - 1)
		for i := 0; i < numRows; i++ {
			if len(s) > 0 {
				if yu == 0 {
					arr[i][j] = s[0]
					s = s[1:]
				} else {
					if i+yu == numRows-1 {
						arr[i][j] = s[0]
						s = s[1:]
						break
					}
				}
			} else {
				break one
			}
		}
	}
	var ans strings.Builder
	for _, v1 := range arr {
		for _, v := range v1 {
			if v != 0 {
				ans.WriteByte(v)
			}
		}
		//ans.Write(v1)
	}

	return ans.String()
}

func convertOfficial2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ret := []byte{}
	n := len(s)
	cycleLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret = append(ret, s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				ret = append(ret, s[j+cycleLen-i])
			}
		}
	}
	return string(ret)
}
