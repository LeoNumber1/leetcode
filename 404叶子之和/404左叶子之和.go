package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 9
	root.Right = new(TreeNode)
	root.Right.Val = 20
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 15
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	fmt.Println(sumOfLeftLeavesBFS(root))
	fmt.Println(sumOfLeftLeavesBFS(nil))
	root1 := new(TreeNode)
	fmt.Println(sumOfLeftLeavesBFS(root1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, isLeft bool)
	dfs = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			ans += node.Val
			return
		}
		dfs(node.Left, true)
		dfs(node.Right, false)
	}

	dfs(root, false)
	return ans
}

func sumOfLeftLeavesBFS(root *TreeNode) int {
	queue := []*TreeNode{root}
	isLeft := []bool{false}
	var ans int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		is := isLeft[0]
		isLeft = isLeft[1:]
		if node == nil {
			continue
		}
		if node.Right == nil && node.Left == nil && is {
			ans += node.Val
			continue
		}
		queue = append(queue, node.Left)
		isLeft = append(isLeft, true)
		queue = append(queue, node.Right)
		isLeft = append(isLeft, false)
	}
	return ans
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func sumOfLeftLeavesOfficialBFS(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if isLeafNode(node.Left) {
				ans += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			q = append(q, node.Right)
		}
	}
	return
}

func dfs(node *TreeNode) (ans int) {
	if node.Left != nil {
		if isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		ans += dfs(node.Right)
	}
	return
}

func sumOfLeftLeavesOfficialDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}
