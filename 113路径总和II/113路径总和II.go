package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//root := new(TreeNode)
	//root.Val = 5
	//root.Left = new(TreeNode)
	//root.Left.Val = 4
	//root.Left.Left = new(TreeNode)
	//root.Left.Left.Val = 11
	//root.Left.Left.Left = new(TreeNode)
	//root.Left.Left.Left.Val = 7
	//root.Left.Left.Right = new(TreeNode)
	//root.Left.Left.Right.Val = 2
	//root.Right = new(TreeNode)
	//root.Right.Val = 8
	//root.Right.Left = new(TreeNode)
	//root.Right.Left.Val = 13
	//root.Right.Right = new(TreeNode)
	//root.Right.Right.Val = 4
	//root.Right.Right.Left = new(TreeNode)
	//root.Right.Right.Left.Val = 5
	//root.Right.Right.Right = new(TreeNode)
	//root.Right.Right.Right.Val = 1

	str := "[5,4,8,11,null,13,4,7,2,null,null,5,1]"
	//str = "[1,null,2]"
	//str = "[1,null,2,null,3]"
	//str = "[1]"
	root := generateTree(str)

	sum := 22
	//sum = 1
	//fmt.Println(pathSumOfficial(root, sum))
	fmt.Println(pathSumOfficialBFS(root, sum))
	//fmt.Println(pathSum(nil, 0))
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//4 ms-93.80%	4.6 MB-51.66%
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans [][]int
		tmp []int
		dfs func(node *TreeNode, s int)
	)

	dfs = func(node *TreeNode, s int) {
		s += node.Val
		tmp = append(tmp, node.Val)

		if isLeaf(node) {
			if s == sum {
				ans = append(ans, append([]int(nil), tmp...))
			}
			tmp = tmp[:len(tmp)-1]
			return
		}

		if node.Left != nil {
			dfs(node.Left, s)
		}
		if node.Right != nil {
			dfs(node.Right, s)
		}
		s -= node.Val
		tmp = tmp[:len(tmp)-1]
	}

	dfs(root, 0)
	return ans
}

func isLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

//0 ms-100.00%	4.7 MB-50.92%
func pathSumOfficial(root *TreeNode, sum int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, sum)
	return
}

type pair struct {
	node *TreeNode
	left int
}

//8 ms-26.45%	5 MB-26.94%
func pathSumOfficialBFS(root *TreeNode, sum int) (ans [][]int) {
	if root == nil {
		return
	}

	parent := map[*TreeNode]*TreeNode{}

	getPath := func(node *TreeNode) (path []int) {
		for ; node != nil; node = parent[node] {
			path = append(path, node.Val)
		}
		for i, j := 0, len(path)-1; i < j; i++ {
			path[i], path[j] = path[j], path[i]
			j--
		}
		return
	}

	queue := []pair{{root, sum}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, left})
			}
		}
	}

	return
}
