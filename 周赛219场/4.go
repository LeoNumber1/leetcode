package main

import (
	"fmt"
	"sort"
)

func main() {
	cuboids := [][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}
	cuboids = [][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}
	cuboids = [][]int{{12, 76, 13}, {68, 55, 30}, {48, 85, 52}, {91, 7, 41}, {29, 65, 35}}
	cuboids = [][]int{{34, 29, 33}, {7, 25, 75}, {31, 15, 68}, {80, 80, 38}, {72, 21, 30}, {2, 66, 25}, {59, 36, 6}, {39, 48, 95}, {35, 85, 71}, {17, 14, 78}}

	fmt.Println(maxHeight(cuboids))
}

func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	if n == 1 {
		ans := cuboids[0][0]
		for i := 0; i < 3; i++ {
			if cuboids[0][i] > ans {
				ans = cuboids[0][i]
			}
		}
		return ans
	}
	for k := range cuboids {
		sort.Slice(cuboids[k], func(i, j int) bool {
			return cuboids[k][i] > cuboids[k][j]
		})
	}
	sort.Slice(cuboids, func(i, j int) bool {
		return cuboids[i][0] > cuboids[j][0] || cuboids[i][0] == cuboids[j][0] && cuboids[i][1] > cuboids[j][1] || cuboids[i][0] == cuboids[j][0] && cuboids[i][1] == cuboids[j][1] && cuboids[i][2] == cuboids[j][2]
	})
	var ans int
	for i := 0; i < n-1; i++ {
		prev := cuboids[i]
		temp := prev[0]
		for j := i + 1; j < n; j++ {
			if cuboids[j][0] <= prev[0] && cuboids[j][1] <= prev[1] && cuboids[j][2] <= prev[2] {
				prev = cuboids[j]
				temp += prev[0]
			}
		}
		if temp > ans {
			ans = temp
		}
	}
	return ans
}
