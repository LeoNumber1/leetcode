package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	type args struct {
		n           int
		connections string
	}

	tests := []struct {
		index  int
		args   args
		target int
	}{
		{1, args{6, "[[0,1],[0,2],[0,3],[1,2],[1,3]]"}, 2},
		{2, args{6, "[[0,1],[0,2],[0,3],[1,2]]"}, -1},
		{3, args{5, "[[0,1],[0,2],[3,4],[2,3]]"}, 0},
		{4, args{4, "[[0,1],[0,2],[1,2]]"}, 1},
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
		result := makeConnected(tt.args.n, str2matrix(tt.args.connections))
		if tt.target != result {
			errNum = true
			fmt.Println("—————— err in index:", tt.index, "except:", tt.target, " result:", result)
		}
	}

	if !errNum {
		fmt.Println("------- All tests OK! -------")
	}
}

func makeConnected(n int, connections [][]int) int {
	line := len(connections)
	if line < n-1 {
		return -1
	}

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, connection := range connections {
		union(connection[0], connection[1])
	}
	var temp int
	for x, y := range parent {
		if x == y {
			temp++
		}
	}

	return temp - 1
}
