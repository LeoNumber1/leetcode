package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "[1,null,2,3]"
	str = "[6,2,8,0,4,7,9,null,null,3,5]"
	root := generateTree(str)

	fmt.Println(postorderTraversalOfficial2(root))
}

//0 ms-100.00%	2 MB-69.00% 递归
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}

//广度优先，迭代1  0 ms-100.00%	2 MB-29.67%
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = []*TreeNode{root}
	var ret []int
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, p.Val)
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}
	n := len(ret)
	for i := 0; i < n/2; i++ {
		ret[i], ret[n-1-i] = ret[n-1-i], ret[i]
	}
	return ret
}

//广度优先，迭代2  0 ms-100.00%	2 MB-60.50%
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, res := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			stack.Push(cur)
			stack.Push(nil) //已经处理完递归，待读取数据标记
		} else {
			res = append(res, stack.Pop().Val)
		}
	}
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}

//广度优先，迭代3  0 ms-100.00%	2 MB-60.50%
func postorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, res := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		res = append(res, cur.Val)
	}
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

//0 ms	2 MB
func postorderTraversalOfficial(root *TreeNode) []int {
	stack := []*TreeNode{}
	var ans []int
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ans
}

func postorderTraversalOfficial2(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		path := []int{}
		for ; node != nil; node = node.Right {
			path = append(path, node.Val)
		}
		for i := len(path) - 1; i >= 0; i-- {
			res = append(res, path[i])
		}
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
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
