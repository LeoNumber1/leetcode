package main

import (
	"fmt"
	"math"
)

func main() {
	a := 6
	b := 10

	fmt.Println("a & b =", a&b)
	fmt.Println("a ^ b =", a^b) //异或，相异为1，相同为0
	fmt.Println("a | b =", a|b)
	fmt.Println("a &^ b =", a&^b) //&^     位清空 (AND NOT)
	//计算x&^y 首先我们先换算成2进制  0000 0010 &^ 0000 0100 = 0000 0010
	//如果ybit位上的数是0则取x上对应位置的值， 如果ybit位上为1则取结果位上取0
	//6和10 0000 0110 &^ 0000 1010 = 0000 0100 -> 4

	fmt.Println("^a =", ^a)
	//所有正整数的按位取反是其本身+1的负数
	//所有负整数的按位取反是其本身+1的绝对值
	//零的按位取反是 -1
	fmt.Println(1<<63 - 1)

	fmt.Println(math.Exp2(64))
	//fmt.Println(math.MaxUint64)
	fmt.Println(math.MaxInt64)

	fmt.Println(a & -a)            //找出最后一位1的二进制数，因为-a是以补码的形式储存
	fmt.Println(8 & -8)            //找出最后一位1的二进制数，因为-a是以补码的形式储存
	fmt.Println(math.Log2(8 & -8)) //最小的一位的下标,3
	//	v&(v-1) 删除最后一位1
}
