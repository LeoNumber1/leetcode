package main

import "fmt"

func main() {
	//	[1,null,2,3]
	root := new(TreeNode)
	//root.Val = 1
	//root.Right = new(TreeNode)
	//root.Right.Val = 2
	//root.Right.Left = new(TreeNode)
	//root.Right.Left.Val = 3

	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4
	root.Right = new(TreeNode)
	root.Right.Val = 5
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 6

	//fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversal3(root))
	//fmt.Println(preorderTraversal2(nil))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	return inBefore(root, res)
}

func inBefore(root *TreeNode, res []int) []int {
	res = append(res, root.Val)
	if root.Left != nil {
		res = inBefore(root.Left, res)
	}
	if root.Right != nil {
		res = inBefore(root.Right, res)
	}
	return res
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append([]int{root.Val}, preorderTraversal1(root.Left)...)
	rest = append(rest, preorderTraversal1(root.Right)...)
	return rest
}

//深度优先，迭代
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, rest := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
			stack.Push(cur)
			stack.Push(nil) // 已处理完递归，待读取数据标记
		} else {
			rest = append(rest, stack.Pop().Val)
		}
	}
	return rest
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

//深度优先，迭代2
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	stack = append(stack, root)

	var ret []int
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		ret = append(ret, p.Val)
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}

	return ret
}
