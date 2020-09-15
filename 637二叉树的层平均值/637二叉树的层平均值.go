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

	fmt.Println(averageOfLevels(root))
	fmt.Println(averageOfLevelsDFS(root))
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

//广度优先	12 ms	6.9 MB
func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{
		root,
	}
	ans := []float64{}
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(n))
		queue = queue[n:]
	}

	return ans
}

//深度优先	8 ms	7 MB
func averageOfLevelsDFS(root *TreeNode) []float64 {
	var ans []float64
	var count []int
	var vals []int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if len(vals) <= depth {
			count = append(count, 1)
			vals = append(vals, root.Val)
		} else {
			count[depth]++
			vals[depth] += root.Val
		}
		if root.Left != nil {
			dfs(root.Left, depth+1)
		}
		if root.Right != nil {
			dfs(root.Right, depth+1)
		}
	}

	dfs(root, 0)

	for i := 0; i < len(vals); i++ {
		ans = append(ans, float64(vals[i])/float64(count[i]))
	}

	return ans
}
