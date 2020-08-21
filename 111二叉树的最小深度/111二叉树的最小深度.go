package main

import (
	"fmt"
	"math"
)

func main() {
	//3,9,20,null,null,15,7
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 9
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 20
	root.Right.Left = new(TreeNode)
	//root.Right.Left.Val = 15
	//root.Right.Right = new(TreeNode)
	//root.Right.Right.Val = 7

	//root.Left.Left.Left = new(TreeNode)
	//root.Left.Left.Left.Val = 2

	fmt.Println(minDepth(root))
	fmt.Println(minDepth1(root))
	fmt.Println(minDepthOfficial1(root))
	fmt.Println(minDepthOfficial2(root))
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

func minDepth(root *TreeNode) int {
	var f func(root *TreeNode, isRoot bool) int

	f = func(root *TreeNode, isRoot bool) int {
		if root == nil {
			if !isRoot {
				//如果不是根节点，
				return math.MaxInt32
			} else {
				return 0
			}
		}
		if root.Left == nil && root.Right == nil {
			return 1 //如果是叶子节点，最小高度是1
		}

		return min(f(root.Left, false), f(root.Right, false)) + 1
	}

	return f(root, true)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth1(root *TreeNode) int {
	var f func(root *TreeNode, isRoot bool) int

	f = func(root *TreeNode, isRoot bool) int {
		if root == nil {
			if !isRoot {
				//如果不是根节点，
				return -1
			} else {
				return 0
			}
		}
		if root.Left == nil && root.Right == nil {
			return 1 //如果是叶子节点，最小高度是1
		}

		return min1(f(root.Left, false), f(root.Right, false)) + 1
	}

	return f(root, true)
}

func min1(a, b int) int {
	if a < 0 {
		return b
	}
	if b < 0 {
		return a
	}

	if a < b {
		return a
	}
	return b
}

//深度优先
func minDepthOfficial1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

//广度优先
func minDepthOfficial2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
