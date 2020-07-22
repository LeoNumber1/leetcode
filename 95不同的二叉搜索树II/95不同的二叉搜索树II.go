package main

import "fmt"

var count = 3
var num int

func main() {
	n := count
	ret := generateTreesOfficial(n)
	for _, v := range ret {
		num = 0
		var res []interface{}
		v.InBegin(&res)
		fmt.Println(res)
	}
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
func (node *TreeNode) InBegin(result *[]interface{}) {
	*result = append(*result, node.Val)
	num++
	if node.Left != nil {
		node.Left.InBegin(result)
	} else {
		if num < count {
			*result = append(*result, "null")
		}
	}
	if node.Right != nil {
		node.Right.InBegin(result)
	} else {
		if num < count {
			*result = append(*result, "null")
		}
	}
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

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	if n <= 1 {
		tn := new(TreeNode)
		tn.Val = n
		return append(res, tn)
	}
	t := new(TreeNode)
	tn := t
	for i := 1; i < n+1; i++ {
		//tn.Right = new(TreeNode)
		//tn := new(TreeNode)
		tn.Val = i
		for j := 1; j < n+1; j++ {
			tnj := new(TreeNode)
			tnj.Val = j

			if i < j {
				tn.Right = tnj
				tn = tn.Right
			} else if i > j {
				tn.Left = tnj
				tn = tn.Left
			}
		}
		res = append(res, t)
	}
	return res
}

func generateTrees2(n int) []*TreeNode {
	var res []*TreeNode
	if n <= 1 {
		tn := new(TreeNode)
		tn.Val = n
		return append(res, tn)
	}
	for i := 1; i < n+1; i++ {
		root := new(TreeNode)
		root.Val = i
		for j := 1; j < n+1; j++ {
			root.Insert(j)
		}
		res = append(res, root)
	}
	return res
}

func (node *TreeNode) Insert(v int) {
	if v < node.Val {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &TreeNode{v, nil, nil}
		}
	} else if v > node.Val {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &TreeNode{v, nil, nil}
		}
	}
}

//panic: interface conversion: *main.TreeNode is not leetcode.OBJInterface: missing method Serialize
//_/leetcode/user_code/interpret_1595312530.3449352_TJ7LfmX12P/precompiled.(*Serializer).Serialize(0xc00006ad00, 0x4d7580, 0xc000076060, 0x1, 0x4d7580)
//serializer.go, line 102
//_/leetcode/user_code/interpret_1595312530.3449352_TJ7LfmX12P/precompiled.(*Serializer).Serialize(0xc00006ae58, 0x4cc520, 0xc0000761a0, 0xc0000761a0, 0x5a71d3)
//serializer.go, line 94
//main.main()
//solution.go, line 88

func generateTreesOfficial(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
