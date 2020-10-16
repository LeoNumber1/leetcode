package main

import (
	"strconv"
	"strings"
)

func main() {
	s := "[1,2,3,4,5,6,7]"
	root := generateTree(s)

	//connect(root)
	connectOfficial(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//8 ms-63.46%	6.5 MB-12.19%
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue = []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			this := queue[0]
			queue = queue[1:]
			if i+1 < n {
				this.Next = queue[0]
			}
			if this.Left != nil {
				queue = append(queue, this.Left)
			}
			if this.Right != nil {
				queue = append(queue, this.Right)
			}
		}
	}
	return root
}

//超时了，QAQ
func connect1(root *Node) *Node {
	var f func(node *Node) *Node
	f = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if node.Left == nil {
			return nil
		}
		node.Left.Next = node.Right
		node.Right.Next = f(node.Next)
		f(node.Left)
		return node.Left
	}
	f(root)
	return root
}

//8 ms	6.3 MB
func connectOfficial1(root *Node) *Node {
	if root == nil {
		return nil
	}
	node := root
	var pre *Node
	for node != nil && node.Left != nil {
		if pre == nil {
			pre = node
		}
		node.Left.Next = node.Right
		if node.Next == nil {
			node = pre.Left
			pre = nil
		} else {
			node.Right.Next = node.Next.Left
			node = node.Next
		}
	}
	return root
}

//4 ms-95.15%	6.3 MB-51.08%
func connectOfficial(root *Node) *Node {
	if root == nil {
		return nil
	}
	//每次循环从该层最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		//通过Next遍历这一层节点，为下一层节点更新Next指针
		for node := leftmost; node != nil; node = node.Next {
			//左节点指向右节点
			node.Left.Next = node.Right
			//右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
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
