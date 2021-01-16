package main

import "fmt"

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}} //2
	//isConnected = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 //1
	//isConnected = [][]int{{1, 1, 1, 0}, {1, 1, 1, 0}, {1, 1, 1, 0}, {0, 0, 0, 1}}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          //2
	//isConnected = [][]int{{1, 1, 0, 0}, {1, 1, 0, 1}, {0, 0, 0, 0}, {0, 1, 0, 1}}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          //2
	//isConnected = [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          //1
	//isConnected = [][]int{{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, {0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}} //8

	fmt.Println(findCircleNum(isConnected))
}

//28 ms	6.6 MB
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	ans := n
	isVisit := make([]bool, n)
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			return
		}
		for j := 0; j < n; j++ {
			if index != j && !isVisit[j] && isConnected[index][j] == 1 {
				isVisit[j] = true
				ans--
				dfs(j)
			}
		}
	}
	for i := 0; i < n; i++ {
		isVisit[i] = true
		dfs(i)
	}
	return ans
}

//24 ms	6.6 MB
func findCircleNumOfficialDFS(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}

//28 ms	6.6 MB
func findCircleNumOfficialBFS(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	for i, v := range vis {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				vis[from] = true
				for to, conn := range isConnected[from] {
					if conn == 1 && !vis[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return
}

//28 ms	6.6 MB
//todo 并查集
func findCircleNumOfficial(isConnected [][]int) (ans int) {
	n := len(isConnected)
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

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return
}
