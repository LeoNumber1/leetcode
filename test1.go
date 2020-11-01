package main

import (
	"fmt"
	"math"
)

func main() {
	test001()
}

const v int = 1978

func test000() {
	var minVal int = math.MaxInt32
	val1 := v
	for m := 2; m < 10; m++ {
		val1 *= v
		//val2 := exp(v, m)
		val2 := val1 * v
		for n := m + 1; n < 10; n++ {
			val2 *= v
			if val1%1000 == val2%1000 {
				if m+n < minVal {
					fmt.Println("m =", m, " n =", n)
					minVal = m + n
				}
			}
		}
	}
	fmt.Println(minVal)
}

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
