package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := "25525511135"
	//s = "0000"
	//s = "127001"
	//s = "19216811"
	//s = "10246578"
	//s = "255255255255"

	//fmt.Println("335" > "255")
	//fmt.Println("35" > "255")
	//fmt.Println("111" > "255")
	//fmt.Println("0" > "255")

	//fmt.Println(restoreIpAddressesOfficial(s))
	//fmt.Println(restoreIpAddresses1(s))
	t1 := time.Now()
	fmt.Println(restoreIpAddresses2(s), "t1 =", time.Since(t1))
	t2 := time.Now()
	fmt.Println(restoreIpAddresses3(s), "t2 =", time.Since(t2))
}

func restoreIpAddresses(s string) []string {
	sLen := len(s)
	k := sLen / 4
	_ = k
	res := make([]string, 0)
	var ret string
	var index int
	for i := 0; i < 4; i++ {
		if i != 3 {
			if s[index] == '0' {
				ret += string(s[index]) + "."
				index++
				continue
			}

			if s[index:k] < "255" && s[index+k:index+k+k] < "255" {

			}
		} else {
			ret += s[index:]
		}

		//if k+1 < len(s) {
		//	ret += s[:k+1] + "."
		//	s = s[k+1:]
		//} else {
		//	ret += s
		//	s = ""
		//}
	}
	res = append(res, ret)
	return res
}

const SEG_COUNT = 4

var (
	ans      []string
	segments []int
)

func restoreIpAddressesOfficial(s string) []string {
	segments = make([]int, SEG_COUNT)
	ans = []string{}
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, segId, segStart int) {
	// 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
	if segId == SEG_COUNT {
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < SEG_COUNT; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != SEG_COUNT-1 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}

	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}
	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		segments[segId] = 0
		dfs(s, segId+1, segStart+1)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfs(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}

func restoreIpAddresses1(s string) []string {
	ret := []string{}
	dfs1(s, &ret, 0, "")
	return ret
}

func dfs1(s string, ret *[]string, dot int, ip string) {
	if s == "" {
		return
	}
	if dot == 3 {
		if len(s) > 3 {
			return
		}
		if valid(s) {
			*ret = append(*ret, ip+s)
		}
		return
	}
	for i, dot := 0, dot+1; i < 3 && i < len(s); i++ {
		if !valid(s[:i+1]) {
			continue
		}
		dfs1(s[i+1:], ret, dot, ip+s[:i+1]+".")
	}
}

func valid(s string) bool {
	switch {
	case len(s) > 3:
		return false
	case len(s) > 1 && s[0] == '0':
		return false
	case len(s) == 3 && s[0] == '2' && s[1] == '5' && s[2] <= '5':
		return true
	case len(s) == 3 && s[0] == '2' && s[1] < '5':
		return true
	case len(s) == 3 && s[0] == '1':
		return true
	case len(s) > 0 && len(s) < 3:
		return true
	}
	return false
}

func judgeNumber(num string) bool {
	if len(num) > 1 && num[0] == '0' {
		return false
	}
	result := 0
	for i := 0; i < len(num); i++ {
		result = 10*result + int(num[i]-'0')
	}
	if result > 255 {
		return false
	}
	return true
}

func restoreIpAddresses2(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	var result []string
	for i := 1; i < 4 && i < len(s)-2; i++ {
		for j := i + 1; j < i+4 && j < len(s)-1; j++ {
			for k := j + 1; k < j+4 && k < len(s); k++ {
				if judgeNumber(s[0:i]) && judgeNumber(s[i:j]) && judgeNumber(s[j:k]) && judgeNumber(s[k:]) {
					result = append(result, strings.Join([]string{s[0:i], s[i:j], s[j:k], s[k:]}, "."))
				}
			}
		}
	}
	return result
}

func restoreIpAddresses3(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	var result []string
	for i := 1; i < 4 && i < len(s)-2; i++ {
		if judgeNumber3(s[0:i]) {
			for j := i + 1; j < i+4 && j < len(s)-1; j++ {
				if judgeNumber3(s[i:j]) {
					for k := j + 1; k < j+4 && k < len(s); k++ {
						if judgeNumber3(s[j:k]) && judgeNumber3(s[k:]) {
							result = append(result, strings.Join([]string{s[0:i], s[i:j], s[j:k], s[k:]}, "."))
						}
					}
				}
			}
		}
	}
	return result
}

func judgeNumber3(num string) bool {
	if len(num) > 1 && num[0] == '0' {
		return false
	}

	if len(num) > 3 {
		return false
	}
	n, _ := strconv.Atoi(num)
	if n > 255 {
		return false
	}
	return true
}
