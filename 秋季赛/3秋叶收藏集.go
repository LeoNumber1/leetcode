package main

import (
	"fmt"
	"strings"
)

func main() {
	leaves := "rrryyyrryyyrr"
	leaves = "rryyyrrrrryyyyyyyyrryyyrr"
	//leaves = "ryr"
	//leaves = "yyryy"
	//leaves = "rrryyyrrr"
	//leaves = "rrrrrrrrrr"
	//leaves = "yyyyy"

	//fmt.Println(minimumOperations(leaves))
	//fmt.Println(minimumOperations1(leaves))
	//fmt.Println(minimumOperations2(leaves))
	fmt.Println(minimumOperations3(leaves))
}

//这个可以，但超出时间限制了
func minimumOperations2(leaves string) int {
	var count int
	var r = []rune(leaves)
	if leaves[0] != 'r' {
		count++
		r[0] = 'r'
	}
	if leaves[len(r)-1] != 'r' {
		count++
		r[len(r)-1] = 'r'
	}

	leaves = strings.Trim(string(r), "r")
	//fmt.Println(leaves)
	n := len(leaves)
	if n == 0 { //中间没黄，翻一次
		return count + 1
	}

	arr := []int{} //奇数为r，偶数为y，值代表数量

	var pre byte
	var countR, countY int
	var totalR int
	for i := 0; i < n; i++ {
		if leaves[i] == 'r' {
			totalR++
		}
		if leaves[i] != pre {
			if leaves[i] == 'r' {
				arr = append(arr, 1)
				countR++
			} else {
				arr = append(arr, 1)
				countY++
			}
			pre = leaves[i]
		} else {
			arr[len(arr)-1]++
		}
	}

	if countR == 0 { //中间没红，不用翻面，返回0
		return count
	}

	var totalY = len(leaves) - totalR
	var count1 int = 100000 //

	for i := 0; i < len(arr); i++ {
		if i%2 == 0 { //偶数，y
			for j := i; j < len(arr); j = j + 2 {
				changeR, notChangeY := 0, 0
				for k := i; k <= j; k++ {
					if k%2 == 1 { //奇数，r
						changeR += arr[k]
					} else { //偶数，y
						notChangeY += arr[k]
					}
				}
				changeTotal := changeR + totalY - notChangeY

				count1 = min(count1, changeTotal)
			}
		}
	}

	return count + count1
}

func minimumOperations3(leaves string) int {
	var count int
	if leaves[0] != 'r' {
		count++
		leaves = "r" + leaves[1:]
	}
	if leaves[len(leaves)-1] != 'r' {
		count++
		leaves = leaves[:len(leaves)-1] + "r"
	}

	leaves = strings.Trim(leaves, "r")
	n := len(leaves)
	if n == 0 { //中间没黄，翻一次
		return count + 1
	}

	arr := []int{} //奇数为r，偶数为y，值代表数量

	var pre byte
	var countR int
	var totalR int
	for i := 0; i < n; i++ {
		if leaves[i] == 'r' {
			totalR++
		}
		if leaves[i] != pre {
			if leaves[i] == 'r' {
				arr = append(arr, 1)
				countR++
			} else {
				arr = append(arr, 1)
			}
			pre = leaves[i]
		} else {
			arr[len(arr)-1]++
		}
	}

	if countR == 0 { //中间没红，不用翻面，返回0
		return count
	}

	var totalY = len(leaves) - totalR
	var count1 int = 100000 //

	for i := 0; i < len(arr); i = i + 2 {
		changeR, notChangeY := 0, 0

		for j := i; j < len(arr); j = j + 2 {
			notChangeY += arr[j]
			if i != j {
				changeR += arr[j-1]
			}
			changeTotal := changeR + totalY - notChangeY

			count1 = min(count1, changeTotal)
		}
	}

	return count + count1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
