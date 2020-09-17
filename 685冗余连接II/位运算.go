package main

import "fmt"

func main() {
	a := 6
	b := 10

	fmt.Println(a & b)
	fmt.Println(a ^ b)
	fmt.Println(a | b)
	fmt.Println(a &^ b) //&^     位清空 (AND NOT)
	//计算x&^y 首先我们先换算成2进制  0000 0010 &^ 0000 0100 = 0000 0010 如果ybit位上的数是0则取x上对应位置的值， 如果ybit位上为1则取结果位上取0
}
