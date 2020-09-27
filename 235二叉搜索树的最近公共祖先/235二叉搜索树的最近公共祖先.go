package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[6,2,8,0,4,7,9,null,null,3,5]"
	root := generateTree(s)
	p := root.Left.Left
	q := root.Right.Right

	//fmt.Println(lowestCommonAncestor(root, p, q))
	fmt.Println(lowestCommonAncestorOfficial1(root, p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//28 ms-44.75%	7.1 MB-40.77%
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == q {
		return q
	}
	var level, level1, level2, count int
	var queue = []*TreeNode{root}
	parent := make(map[*TreeNode]*TreeNode)
	for len(queue) > 0 {
		level++
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == p {
				level1 = level
				count++
			}
			if node == q {
				level2 = level
				count++
			}
			if count == 2 {
				goto next
			}
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, node.Right)
			}
		}
	}

next:
	if level1 > level2 {
		level1, level2 = level2, level1
		p, q = q, p
	}

	for level1 < level2 {
		q = parent[q]
		level2--
	}

	for p != q {
		p = parent[p]
		q = parent[q]
	}

	return p
}

//24 ms-79.91%	7.7 MB-15.39%
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == q {
		return q
	}

	pathP := []*TreeNode{root}
	node := root
	for node != p {
		if node.Val > p.Val {
			node = node.Left
		} else {
			node = node.Right
		}
		pathP = append(pathP, node)
	}
	pathQ := []*TreeNode{root}
	node = root
	for node != q {
		if node.Val > q.Val {
			node = node.Left
		} else {
			node = node.Right
		}
		pathQ = append(pathQ, node)
	}

	var ans *TreeNode
	for i := 0; i < len(pathQ) && i < len(pathQ) && pathQ[i] == pathP[i]; i++ {
		ans = pathP[i]
	}

	return ans
}

func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}

//28 ms-44.75%	7.7 MB-32.57%
func lowestCommonAncestorOfficial(root, p, q *TreeNode) (ancestor *TreeNode) {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return
}

//24 ms-79.91%	7.7 MB-5.39%
func lowestCommonAncestorOfficial1(root, p, q *TreeNode) (ancestor *TreeNode) {
	var ans *TreeNode
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return ans
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
