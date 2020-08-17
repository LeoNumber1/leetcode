package main

import (
	"fmt"
)

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
	root.Right.Right.Left = new(TreeNode)
	root.Right.Right.Left.Val = 2

	fmt.Println(isBalanced(root))
	fmt.Println(isBalancedOfficial2(root))
	//fmt.Println(height(root))
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func isBalancedOfficial2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalancedOfficial2(root.Left) && isBalancedOfficial2(root.Right) && abs(height(root.Left)-height(root.Right)) <= 1
}
