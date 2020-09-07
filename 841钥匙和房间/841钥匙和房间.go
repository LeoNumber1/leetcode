package main

import "fmt"

func main() {
	rooms := [][]int{
		//{1}, {2}, {3}, {},
		{1, 3}, {3, 0, 1}, {2}, {0},
	}

	fmt.Println(canVisitAllRooms(rooms))
}

//BFS 广度优先
func canVisitAllRooms(rooms [][]int) bool {
	m := len(rooms)
	if m <= 1 {
		return true
	}

	matrix := make([]bool, m)
	matrix[0] = true

	queue := []int{}
	for _, v := range rooms[0] {
		queue = append(queue, v)
	}
	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]
		matrix[room] = true
		for _, v := range rooms[room] {
			if !matrix[v] {
				queue = append(queue, v)
			}
		}
	}
	for i := 0; i < m; i++ {
		if !matrix[i] {
			return false
		}
	}
	return true
}
