package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 1
	//root.Right = new(TreeNode)
	//root.Right.Val = 2
	//root.Right.Left = new(TreeNode)
	//root.Right.Left.Val = 3

	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 5
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 6

	fmt.Println(inorderTraversal(root))
	//fmt.Println(inorderTraversalOfficial(root))
	fmt.Println(inorderTraversalOfficial1(root))
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

//0 ms	2.1 MB
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

//0 ms	2 MB	迭代，栈
func inorderTraversalOfficial(root *TreeNode) []int {
	var ans []int
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}

	return ans
}

//Morris 中序遍历	0 ms	2 MB
func inorderTraversalOfficial1(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			//predecessor节点表示当前root节点向左走一步，然后一直向右走直至无法走的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				//有右子树且没有设置过指向root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				//将predecessor的右指针指向root，这样后面遍历完root.left后，就能通过这个指向回到root
				predecessor.Right = root
				//遍历左子树
				root = root.Left
			} else { //predecessor的右指针已经指向了root，则表示左子树root.Left已经访问完了
				res = append(res, root.Val)
				//复原
				predecessor.Right = nil
				//遍历右子树
				root = root.Right
			}
		} else { //没有左子树
			res = append(res, root.Val)
			//若有右子树，则遍历右子树
			//若没有右子树，则整颗左子树已遍历完，root会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}
