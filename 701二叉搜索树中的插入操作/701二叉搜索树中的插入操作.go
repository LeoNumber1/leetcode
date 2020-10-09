package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[4,2,7,1,3]"
	val := 5

	root := generateTree(s)

	fmt.Println(insertIntoBST1(root, val))
}

//36 ms-78.59%	7.4 MB-6.20%
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var prev *TreeNode
	var isLeft bool
	var f func(treeNode *TreeNode)
	f = func(treeNode *TreeNode) {
		if treeNode == nil {
			if isLeft {
				prev.Left = &TreeNode{Val: val}
			} else {
				prev.Right = &TreeNode{Val: val}
			}
			return
		}
		prev = treeNode
		if treeNode.Val > val {
			isLeft = true
			f(treeNode.Left)
		} else {
			isLeft = false
			f(treeNode.Right)
		}
	}
	f(node)

	return root
}

//44 ms-11.98%	7.4 MB-23.36%
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var prev *TreeNode
	var isLeft bool
	for node != nil {
		prev = node
		if node.Val > val {
			isLeft = true
			node = node.Left
		} else {
			isLeft = false
			node = node.Right
		}
	}

	if isLeft {
		prev.Left = &TreeNode{Val: val}
	} else {
		prev.Right = &TreeNode{Val: val}
	}

	return root
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
