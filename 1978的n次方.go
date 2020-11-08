package main

import (
	"fmt"
)

func main() {
	test001()
}

const v int = 1978

func test001() {
	var m = make(map[int]int)
	var num = v
	for i := 2; i < 1000; i++ {
		num *= v
		tmp := num % 1000
		num %= 100000000
		if val, ok := m[tmp]; !ok {
			m[tmp] = i
		} else {
			fmt.Println("m =", val, " n =", i)
			fmt.Println("相等的后三位是：", tmp)
			return
		}
	}
}
