package main

import "fmt"

func main() {
	moves := "UD"
	moves = "LLRR"

	fmt.Println(judgeCircle(moves))
	fmt.Println(judgeCircle1(moves))
}

func judgeCircle(moves string) bool {
	var numU, numD, numL, numR int
	for _, v := range moves {
		switch v {
		case 'U':
			numU++
		case 'D':
			numD++
		case 'L':
			numL++
		case 'R':
			numR++
		default:
			return false
		}
	}
	return numU == numD && numL == numR
}

func judgeCircle1(moves string) bool {
	start := []int{0, 0}
	for _, v := range moves {
		switch v {
		case 'U':
			start[0]--
		case 'D':
			start[0]++
		case 'L':
			start[1]--
		case 'R':
			start[1]++
		default:
			return false
		}
	}
	return start[0] == 0 && start[1] == 0
}
