package main

import "fmt"

func main() {
	image := [][]int{
		//{1, 1, 1}, {1, 1, 0}, {1, 0, 1},
		{0, 0, 0}, {0, 1, 1},
	}
	sr := 1
	sc := 1
	newColor := 2
	newColor = 1
	fmt.Println(floodFill(image, sr, sc, newColor))
	fmt.Println(floodFill1(image, sr, sc, newColor))
}

var direction = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

//深度优先算法
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	tmp := image[sr][sc]
	image[sr][sc] = newColor
	for _, v := range direction {
		newRow := sr + v[0]
		newCol := sc + v[1]
		if newRow < len(image) && newCol < len(image[sr]) && newRow >= 0 && newCol >= 0 && image[newRow][newCol] == tmp && newColor != tmp {
			floodFill(image, newRow, newCol, newColor)
		}
	}
	return image
}

//广度优先算法
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	tmp := image[sr][sc]
	queue := [][]int{{sr, sc}}
	for len(queue) > 0 {
		src := queue[0]
		queue = queue[1:]
		image[src[0]][src[1]] = newColor
		for _, v := range direction {
			newRow := src[0] + v[0]
			newCol := src[1] + v[1]
			if newRow < len(image) && newCol < len(image[sr]) && newRow >= 0 && newCol >= 0 && image[newRow][newCol] == tmp && newColor != tmp {
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}

	return image
}
