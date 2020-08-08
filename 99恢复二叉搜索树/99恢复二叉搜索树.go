package main

import "fmt"

func main() {
	//1,3,null,null,2
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2

	//3,1,4,null,null,2
	//root.Val = 3
	//root.Left = new(TreeNode)
	//root.Left.Val = 1
	//root.Right = new(TreeNode)
	//root.Right.Val = 4
	//root.Right.Left = new(TreeNode)
	//root.Right.Left.Val = 2

	recoverTreeOfficial2(root)

	fmt.Println(root.Val)
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

func recoverTree(root *TreeNode) {
	nums := []int{}
	root.InOrder(&nums)
	fmt.Println(nums)
	return
}

//中序遍历
func (r *TreeNode) InOrder(nums *[]int) {
	if r.Left != nil {
		r.Left.InOrder(nums)
	}
	*nums = append(*nums, r.Val)
	if r.Right != nil {
		r.Right.InOrder(nums)
	}
}

func recoverTreeOfficial(root *TreeNode) {
	nums := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwapped(nums)
	fmt.Println(x, y)
	recover1(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover1(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover1(root.Right, count, x, y)
	recover1(root.Left, count, x, y)
}

func recoverTreeOfficial2(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}
