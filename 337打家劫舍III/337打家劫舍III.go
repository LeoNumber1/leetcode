package main

import "fmt"

func main() {
	//[3,2,3,null,3,null,1]
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 1

	//fmt.Println(rob(root))
	fmt.Println(robOfficial(root))
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

//8 ms	6.5 MB
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	f := map[*TreeNode]int{} //选择该节点的最大值
	g := map[*TreeNode]int{} //不选择该节点的最大值
	var dfs = func(root *TreeNode) {}
	dfs = func(root *TreeNode) {
		l := root.Left
		r := root.Right
		f[root] = root.Val
		g[root] = 0
		if l != nil {
			dfs(l)
			f[root] += g[l]
			g[root] += max(f[l], g[l])
		}
		if r != nil {
			dfs(r)
			f[root] += g[r]
			g[root] += max(f[r], g[r])
		}
	}

	dfs(root)
	return max(f[root], g[root])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//8 ms	5.9 MB
func robOfficial(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
