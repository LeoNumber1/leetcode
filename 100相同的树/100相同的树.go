package main

import "fmt"

func main() {
	p := new(TreeNode)
	p.Val = 1
	p.Left = new(TreeNode)
	p.Left.Val = 2
	p.Right = new(TreeNode)
	p.Right.Val = 3

	q := new(TreeNode)
	q.Val = 1
	q.Left = new(TreeNode)
	q.Left.Val = 2
	q.Right = new(TreeNode)
	q.Right.Val = 3

	//fmt.Println(isSameTree(p, q))
	fmt.Println(isSameTreeOfficial2(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}

func isSameTreeOfficial1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//广度优先搜索
func isSameTreeOfficial2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1.Val != node2.Val {
			return false
		}
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}
		if left1 != nil {
			queue1 = append(queue1, left1)
		}

		if right1 != nil {
			queue1 = append(queue1, right1)
		}

		if left2 != nil {
			queue2 = append(queue2, left2)
		}

		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}
	return len(queue1) == 0 && len(queue2) == 0
}
