package main

import "fmt"

func main() {
	//[3,9,20,null,null,15,7]
	root := new(TreeNode)
	//root.Val = 3
	//root.Left = new(TreeNode)
	//root.Left.Val = 9
	//root.Right = new(TreeNode)
	//root.Right.Val = 20
	//root.Right.Left = new(TreeNode)
	//root.Right.Left.Val = 15
	//root.Right.Right = new(TreeNode)
	//root.Right.Right.Val = 7

	//[1,2,null,3,null,4,null,5]
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 4
	root.Left.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Left.Val = 5

	var res = new([]int)
	root.InBegin(res)
	fmt.Println(*res)
	fmt.Println(maxDepth(root))
	//fmt.Println(maxDepthOfficial1(root))
	//fmt.Println(maxDepthOfficial2(root))
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tn := root
	depthLeft := 1
	depthRight := 1
	if tn.Left != nil {
		depthLeft++
		depthLeft = max(depthLeft, maxDepth(tn.Left)+1)
		//tn = tn.Left
	}
	if root.Right != nil {
		depthRight++
		depthRight = max(depthRight, maxDepth(root.Right)+1)
		//root = root.Right
	}
	return max(depthLeft, depthRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//中序遍历
func (node *TreeNode) InOrder(result *[]int) {
	if node.Left != nil {
		node.Left.InOrder(result)
	}
	*result = append(*result, node.Val)
	if node.Right != nil {
		node.Right.InOrder(result)
	}
}

//先序遍历
func (node *TreeNode) InBegin(result *[]int) {
	*result = append(*result, node.Val)
	if node.Left != nil {
		node.Left.InBegin(result)
	}
	if node.Right != nil {
		node.Right.InBegin(result)
	}
}

func maxDepthOfficial1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthOfficial1(root.Left), maxDepthOfficial1(root.Right)) + 1
}

//广度优先搜索
func maxDepthOfficial2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
