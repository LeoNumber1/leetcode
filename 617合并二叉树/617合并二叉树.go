package main

import "fmt"

func main() {
	t1 := new(TreeNode)
	t1.Val = 1
	t1.Left = new(TreeNode)
	t1.Left.Val = 3
	t1.Right = new(TreeNode)
	t1.Right.Val = 2
	t1.Left.Left = new(TreeNode)
	t1.Left.Left.Val = 5

	t2 := new(TreeNode)
	t2.Val = 2
	t2.Left = new(TreeNode)
	t2.Left.Val = 1
	t2.Right = new(TreeNode)
	t2.Right.Val = 3
	t2.Right.Right = new(TreeNode)
	t2.Right.Right.Val = 7
	t2.Left.Right = new(TreeNode)
	t2.Left.Right.Val = 4

	t1Ans := []int{}
	//inPre(t1, &t1Ans)

	//a := mergeTreesDfs1(t1, t2)
	a := mergeTrees1(t1, t2)

	//inPre(mergeTrees1(t1, t2), &t1Ans)
	//inPre(mergeTreesBfs(t1, t2), &t1Ans)
	t1.Left.Left.Val = 999
	t2.Left.Right.Val = 111
	inPre(a, &t1Ans)
	fmt.Println(t1Ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//24 ms-45.93%	6.8 MB-39.50%
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
	} else if t2 != nil {
		t1 = new(TreeNode)
		t1.Val = t2.Val
	} else {
		t2 = new(TreeNode)
	}

	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

//24 ms-45.93%	6.9 MB-6.72%	这个才是真正的新二叉树，和原来的树没有关系
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	node := new(TreeNode)
	if t1 != nil && t2 != nil {
		node.Val = t1.Val + t2.Val
	} else if t2 != nil {
		t1 = new(TreeNode)
		node.Val = t2.Val
	} else {
		t2 = new(TreeNode)
		node.Val = t1.Val
	}

	node.Left = mergeTrees1(t1.Left, t2.Left)
	node.Right = mergeTrees1(t1.Right, t2.Right)

	return node
}

//16 ms	6.8 MB
func mergeTreesDfs1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

//20 ms	6.9 MB
func mergeTreesBfs(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	merged := &TreeNode{Val: t1.Val + t2.Val}
	queue := []*TreeNode{merged}
	queue1, queue2 := []*TreeNode{t1}, []*TreeNode{t2}

	for len(queue1) > 0 && len(queue2) > 0 {
		node := queue[0]
		queue = queue[1:]
		node1 := queue1[0]
		queue1 = queue1[1:]
		node2 := queue2[0]
		queue2 = queue2[1:]
		l1, r1 := node1.Left, node1.Right
		l2, r2 := node2.Left, node2.Right
		if l1 == nil {
			node.Left = l2
		} else if l2 == nil {
			node.Left = l1
		} else {
			node.Left = &TreeNode{Val: l1.Val + l2.Val}
			queue = append(queue, node.Left)
			queue1 = append(queue1, node1.Left)
			queue2 = append(queue2, node2.Left)
		}

		if r1 == nil {
			node.Right = r2
		} else if r2 == nil {
			node.Right = r1
		} else {
			node.Right = &TreeNode{Val: r1.Val + r2.Val}
			queue = append(queue, node.Right)
			queue1 = append(queue1, node1.Right)
			queue2 = append(queue2, node2.Right)
		}
	}

	return merged
}

func inPre(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	inPre(root.Left, ret)
	inPre(root.Right, ret)
}
