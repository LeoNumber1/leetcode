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
		target string
	}{
		{1, args{1, "[[2,1],[3,4],[3,2]]"}, "[1 2 3 4]"},
		{2, args{1, "[[1,2],[3,2],[3,4]]"}, "[1 2 3 4]"},
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
		result := restoreArrayOfficial(str2matrix(tt.args.connections))
		target := fmt.Sprint(result)
		if tt.target != target {
			errNum = true
			fmt.Println("—————— err in index:", tt.index, "except:", tt.target, " get result:", result)
		}
	}

	if !errNum {
		fmt.Println("------- All tests are OK! -------")
	}

}

// 超时了，QAQ
func restoreArray0(adjacentPairs [][]int) []int {
	m := make(map[int]struct{})
	arr := make([][]int, 0)
	for _, pair := range adjacentPairs {
		_, has0 := m[pair[0]]
		_, has1 := m[pair[1]]
		if !has0 && !has1 {
			arr = append(arr, pair)
			m[pair[0]] = struct{}{}
			m[pair[1]] = struct{}{}
		} else {
			if has0 && has1 {
				var k0, k1 int = -1, -1

				for k, v := range arr {
					if v[0] == pair[0] {
						k0 = k
						if k1 != -1 {
							break
						}
					} else if v[len(v)-1] == pair[0] {
						k0 = k
						if k1 != -1 {
							break
						}
					}
					if v[0] == pair[1] {
						k1 = k
						if k0 != -1 {
							break
						}
					} else if v[len(v)-1] == pair[1] {
						k1 = k
						if k0 != -1 {
							break
						}
					}
				}

				if arr[k0][0] == pair[0] {
					if arr[k1][0] == pair[1] {
						arr[k0] = append(reverse(arr[k1]), arr[k0]...)
					} else {
						arr[k0] = append(arr[k1], arr[k0]...)
					}
				} else {
					if arr[k1][0] == pair[1] {
						arr[k0] = append(arr[k0], arr[k1]...)
					} else {
						arr[k0] = append(arr[k0], reverse(arr[k1])...)
					}
				}
				arr = append(arr[:k1], arr[k1+1:]...)
				delete(m, pair[0])
				delete(m, pair[1])
			} else if has0 {
				for k, v := range arr {
					if v[0] == pair[0] {
						arr[k] = append([]int{pair[1]}, arr[k]...)
						break
					} else if v[len(v)-1] == pair[0] {
						arr[k] = append(arr[k], pair[1])
						break
					}
				}
				m[pair[1]] = struct{}{}
				delete(m, pair[0])
			} else {
				for k, v := range arr {
					if v[0] == pair[1] {
						arr[k] = append([]int{pair[0]}, arr[k]...)
						break
					} else if v[len(v)-1] == pair[1] {
						arr[k] = append(arr[k], pair[0])
						break
					}
				}
				m[pair[0]] = struct{}{}
				delete(m, pair[1])
			}
		}
	}
	return arr[0]
}

//204 ms    37.7 MB
func restoreArray(adjacentPairs [][]int) []int {
	m := make(map[int][]int)
	for _, pair := range adjacentPairs {
		v0, has0 := m[pair[0]]
		v1, has1 := m[pair[1]]
		if !has0 && !has1 {
			m[pair[0]] = pair
			m[pair[1]] = pair
		} else {
			if has0 && has1 {
				var find []int
				if v0[0] == pair[0] {
					if v1[0] == pair[1] {
						find = append(reverse(v1), v0...)
					} else {
						find = append(v1, v0...)
					}
				} else {
					if v1[0] == pair[1] {
						find = append(v0, v1...)
					} else {
						find = append(v0, reverse(v1)...)
					}
				}
				m[find[0]] = find
				m[find[len(find)-1]] = find
				delete(m, pair[0])
				delete(m, pair[1])
			} else if has0 {
				var find []int
				if v0[0] == pair[0] {
					find = v0
					find = append([]int{pair[1]}, find...)
					//找到数组的对应一边，赋值
					m[find[len(find)-1]] = find
				} else {
					find = v0
					find = append(find, pair[1])
					//找到数组的对应一边，赋值
					m[find[0]] = find
				}
				m[pair[1]] = find
				delete(m, pair[0])
			} else {
				// pair1 find
				var find []int
				if m[pair[1]][0] == pair[1] {
					find = v1
					find = append([]int{pair[0]}, find...)
					m[find[len(find)-1]] = find
				} else {
					find = v1
					find = append(find, pair[0])
					m[find[0]] = find
				}
				m[pair[0]] = find
				delete(m, pair[1])
			}
		}
	}
	var ans []int
	for _, v := range m {
		ans = v
		break
	}
	return ans
}

func reverse(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}

//188 ms    27.6 MB
func restoreArrayOfficial(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	g := make(map[int][]int, n)
	for _, p := range adjacentPairs {
		v, w := p[0], p[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	for e, adj := range g {
		if len(adj) == 1 {
			ans[0] = e
			break
		}
	}

	ans[1] = g[ans[0]][0]
	for i := 2; i < n; i++ {
		adj := g[ans[i-1]]
		if ans[i-2] == adj[0] {
			ans[i] = adj[1]
		} else {
			ans[i] = adj[0]
		}
	}
	return ans
}
