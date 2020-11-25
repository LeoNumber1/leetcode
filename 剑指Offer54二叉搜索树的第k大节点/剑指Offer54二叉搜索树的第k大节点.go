package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "3,1,4,null,2"
	k := 1

	root := generateTree(s)
	fmt.Println(kthLargest(root, k))
}

//12 ms	6.4 MB
func kthLargest(root *TreeNode, k int) int {
	var arr []int
	var inOrder func(node *TreeNode, k int)
	inOrder = func(node *TreeNode, k int) {
		if node == nil {
			return
		}
		inOrder(node.Right, k)
		arr = append(arr, node.Val)
		k--
		if k == 0 {
			return
		}
		inOrder(node.Left, k)
	}
	inOrder(root, k)
	return arr[k-1]
}

//16 ms	6.2 MB
func kthLargest0(root *TreeNode, k int) int {
	var arr []int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		arr = append(arr, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	if k > len(arr) {
		return 0
	}
	return arr[len(arr)-k]
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
