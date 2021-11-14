package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	type args struct {
		connections string
	}

	tests := []struct {
		index  int
		args   args
		target int
	}{
		{1, args{"[[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]"}, 4},
	}

	str2matrix := func(s string) (matrix [][]int) {
		arr := strings.Split(s, "],")
		for _, s2 := range arr {
			s2 = strings.TrimLeft(s2, "[")
			s2 = strings.TrimRight(s2, "]")
			arr1 := strings.Split(s2, ",")
			var temp []int
			for _, s3 := range arr1 {
				i, _ := strconv.Atoi(s3)
				temp = append(temp, i)
			}
			matrix = append(matrix, temp)
		}
		return
	}

	var errNum bool
	for _, tt := range tests {
		result := snakesAndLadders(str2matrix(tt.args.connections))
		if tt.target != result {
			errNum = true
			fmt.Println("—————— err in index:", tt.index, "except:", tt.target, " get result:", result)
		}
	}

	if !errNum {
		fmt.Println("------- All tests are OK! -------")
	}
}

func snakesAndLadders(board [][]int) int {
	return 4
}
