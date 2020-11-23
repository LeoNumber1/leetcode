package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}

	//fmt.Println(findMinArrowShots(points))
	fmt.Println(findMinArrowShotsOfficial(points))
}

//、92 ms	7.6 MB
func findMinArrowShots0(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	var arr = make([][]int, 0, n)
	arr = append(arr, points[0])
	for i := 1; i < n; i++ {
		balloon := arr[len(arr)-1]
		if balloon[1] >= points[i][1] {
			arr[len(arr)-1] = points[i]
		} else if points[i][0] <= balloon[1] {
			arr[len(arr)-1] = []int{points[i][0], balloon[1]}
		} else {
			arr = append(arr, points[i])
		}
	}
	return len(arr)
}

//92 ms-32.82%	7.3 MB-46.15%
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	ans := n
	balloon := points[0]
	for i := 1; i < n; i++ {
		if points[i][1] <= balloon[1] {
			ans--
			balloon = points[i]
		} else if points[i][0] <= balloon[1] {
			ans--
			balloon = []int{points[i][0], balloon[1]}
		} else {
			balloon = points[i]
		}
	}
	return ans
}

//88 ms-60.31%	7.4 MB-23.85%
func findMinArrowShotsOfficial(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}
