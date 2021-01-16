package main

import "fmt"

func main() {
	edges := [][]int{{3, 4}, {1, 2}, {2, 4}, {2, 5}, {3, 5}}

	//fmt.Println(findRedundantConnection(edges))
	fmt.Println(findRedundantConnectionOfficial(edges))
}

//20 ms	4.2 MB
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	colors := make(map[int]map[int]bool)
	numMap := make([]int, n+1)
	var ans []int
	for _, edge := range edges {
		if numMap[edge[0]] == 0 && numMap[edge[1]] == 0 {
			for i := 1; i <= n/2; i++ {
				if colors[i] != nil {
					if colors[i][edge[0]] || colors[i][edge[1]] {
						numMap[edge[0]] = i
						numMap[edge[1]] = i
						colors[i][edge[0]] = true
						colors[i][edge[1]] = true
						break
					}
				} else {
					colors[i] = make(map[int]bool)
					colors[i][edge[0]] = true
					colors[i][edge[1]] = true
					numMap[edge[0]] = i
					numMap[edge[1]] = i
					break
				}
			}
		} else if numMap[edge[0]] != 0 && numMap[edge[1]] != 0 {
			if numMap[edge[0]] == numMap[edge[1]] {
				ans = edge
				break
			} else if numMap[edge[0]] < numMap[edge[1]] {
				temp := numMap[edge[1]]
				for key := range colors[temp] {
					numMap[key] = numMap[edge[0]]
					delete(colors[temp], key)
					colors[numMap[edge[0]]][key] = true
				}
			} else {
				temp := numMap[edge[0]]
				for key := range colors[temp] {
					numMap[key] = numMap[edge[1]]
					delete(colors[temp], key)
					colors[numMap[edge[1]]][key] = true
				}
			}
		} else if numMap[edge[0]] != 0 {
			numMap[edge[1]] = numMap[edge[0]]
			colors[numMap[edge[0]]][edge[1]] = true
		} else {
			numMap[edge[0]] = numMap[edge[1]]
			colors[numMap[edge[1]]][edge[0]] = true
		}
	}

	return ans
}

//4 ms	3.1 MB
func findRedundantConnectionOfficial(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
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
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
