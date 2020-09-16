package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 4
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 7
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3

	ans := []int{}
	inPre(root, &ans)
	fmt.Println(ans)

	arr := []int{}
	inPre(invertTree(root), &arr)
	fmt.Println(arr)
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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

//先序遍历
func inPre(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	inPre(root.Left, ans)
	inPre(root.Right, ans)
	return
}
