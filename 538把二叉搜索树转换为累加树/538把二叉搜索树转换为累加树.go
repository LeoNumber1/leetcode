package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Right = new(TreeNode)
	root.Right.Val = 13
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 8

	fmt.Println(convertBST(root))
	//fmt.Println(convertBSTOfficial(root))
	//fmt.Println(convertBSTOfficial1(root))
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

/*
执行用时：16 ms-78.89%
内存消耗：7 MB-25.60%
*/
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, num int) int {
		//num代表父节点传下来的值
		//每个节点只需要加num和右子节点最终值
		if root.Right == nil { //右子节点为空
			root.Val += num
		} else {
			//右子节点不为空
			root.Val += dfs(root.Right, num)
		}

		if root.Left == nil {
			return root.Val
		}

		//左子节点不为空
		return dfs(root.Left, root.Val)
	}
	dfs(root, 0)
	return root
}

//反序中序遍历20 ms-30.74%	7 MB-25.60%
func convertBSTOfficial(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

//Morris 遍历	20 ms-30.74%	7 MB-13.69%
func convertBSTOfficial1(root *TreeNode) *TreeNode {
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}
