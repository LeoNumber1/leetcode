package main

import (
	"fmt"
	"math/big"
)

func main() {
	//var i int64
	//for i = 1; i <= 40; i++ {
	//	fmt.Println(a(big.NewInt(i)))
	//}
	fmt.Println(a(big.NewInt(15)))
}

func a(s *big.Int) *big.Int {
	if s.Int64() == 1 {
		return big.NewInt(1)
	} else {
		return s.Mul(s, a(big.NewInt(s.Int64()-1)))
	}
}
