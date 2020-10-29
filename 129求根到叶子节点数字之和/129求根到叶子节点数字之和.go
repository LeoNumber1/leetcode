package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[4,9,0,5,1]"
	root := generateTree(s)
	fmt.Println(sumNumbers(root))
	fmt.Println(sumNumbersBFS(root))
}

//0 ms-100.00%	2.4 MB-45.59%
func sumNumbers(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, i int)
	dfs = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		i = i*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += i
			return
		}
		dfs(node.Left, i)
		dfs(node.Right, i)
	}
	dfs(root, 0)
	return ans
}

//0 ms-100.00%	2.5 MB-13.03%
func sumNumbersBFS(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}
	nodeQueue := []*TreeNode{root}
	numQueue := []int{root.Val}
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		num := numQueue[0]
		numQueue = numQueue[1:]

		if node.Left == nil && node.Right == nil {
			ans += num
			continue
		}
		num *= 10
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			numQueue = append(numQueue, num+node.Left.Val)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			numQueue = append(numQueue, num+node.Right.Val)
		}
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree(str string) (root *TreeNode) {
	s := strings.TrimLeft(str, "[")
	s = strings.TrimRight(s, "]")
	arr := strings.Split(s, ",")
	if len(arr) == 0 || arr[0] == "null" {
		return
	}
	root = new(TreeNode)
	root.Val, _ = strconv.Atoi(arr[0])
	arr = arr[1:]
	queue := []*TreeNode{root}
	for len(queue) > 0 && len(arr) > 0 {
		node := queue[0]
		queue = queue[1:]

		if arr[0] != "null" {
			node.Left = new(TreeNode)
			node.Left.Val, _ = strconv.Atoi(arr[0])
			queue = append(queue, node.Left)
		}
		arr = arr[1:]
		if len(arr) > 0 {
			if arr[0] != "null" {
				node.Right = new(TreeNode)
				node.Right.Val, _ = strconv.Atoi(arr[0])
				queue = append(queue, node.Right)
			}
			arr = arr[1:]
		}
	}
	return
}
