package main

func main() {

}

//空间复杂度O(mn)版本

func setZeroesMN(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	queue := make([][2]int, 0)
	for i, v := range matrix {
		for j, val := range v {
			if val == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		data := queue[0]
		queue = queue[1:]
		for i := 0; i < m; i++ {
			matrix[i][data[1]] = 0
		}
		for j := 0; j < n; j++ {
			matrix[data[0]][j] = 0
		}
	}
}

//空间复杂度O(m+n)版本
func setZeroesMPlusN(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	mapi := make(map[int]struct{})
	mapj := make(map[int]struct{})
	for i, v := range matrix {
		for j, val := range v {
			if val == 0 {
				mapi[i] = struct{}{}
				mapj[j] = struct{}{}
			}
		}
	}
	for i := range mapi {
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}
	for j := range mapj {
		for i := 0; i < m; i++ {
			matrix[i][j] = 0
		}
	}
}
