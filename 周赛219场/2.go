package main

import "fmt"

func main() {
	n := "32"

	fmt.Println(minPartitions(n))
}

func minPartitions(n string) int {
	var ans int
	for _, v := range n {
		if ans < int(v-'0') {
			ans = int(v - '0')
		}
	}
	return ans
}
