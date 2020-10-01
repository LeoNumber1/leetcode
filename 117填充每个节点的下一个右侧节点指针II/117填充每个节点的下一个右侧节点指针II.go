package main

import (
	"strconv"
	"strings"
)

func main() {
	str := "[1,2,3,4,5,null,7]"
	root := generateTree(str)

	connect(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//4 ms-76.10%	6.2 MB-10.90%
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i+1 < n {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

//4 ms	6.1 MB
func connectOfficial(root *Node) *Node {
	node := root
	for node != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := node; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		node = nextStart
	}
	return root
}

func generateTree(str string) (root *Node) {
	s := strings.TrimLeft(str, "[")
	s = strings.TrimRight(s, "]")
	arr := strings.Split(s, ",")
	if len(arr) == 0 || arr[0] == "null" {
		return
	}
	root = new(Node)
	root.Val, _ = strconv.Atoi(arr[0])
	arr = arr[1:]
	queue := []*Node{root}
	for len(queue) > 0 && len(arr) > 0 {
		node := queue[0]
		queue = queue[1:]

		if arr[0] != "null" {
			node.Left = new(Node)
			node.Left.Val, _ = strconv.Atoi(arr[0])
			queue = append(queue, node.Left)
		}
		arr = arr[1:]
		if len(arr) > 0 {
			if arr[0] != "null" {
				node.Right = new(Node)
				node.Right.Val, _ = strconv.Atoi(arr[0])
				queue = append(queue, node.Right)
			}
			arr = arr[1:]
		}
	}
	return
}
