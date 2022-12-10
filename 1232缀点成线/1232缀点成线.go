package main

import "fmt"

func main() {
	coordinates := [][]int{{2, 1}, {4, 2}, {6, 3}}

	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) bool {
	x1, y1 := float32(coordinates[0][0]), float32(coordinates[0][1])
	x2, y2 := float32(coordinates[1][0]), float32(coordinates[1][1])
	if x1 == x2 {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != int(x1) {
				return false
			}
		}
		return true
	}
	a := (y2 - y1) / (x2 - x1)
	var b = y1 - a*x1
	f := func(x int) int {
		return int(a*float32(x) + b)
	}
	for i := 2; i < len(coordinates); i++ {
		if coordinates[i][1] != f(coordinates[i][0]) {
			return false
		}
	}
	return true
}
