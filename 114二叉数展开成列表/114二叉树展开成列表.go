package main

func main() {
	/*
		    1
		   / \
		  2   5
		 / \   \
		3   4   6
	*/
	root := new(TreeNode)
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
	//flatten(root)
	flattenOfficial(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattenOfficial(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func flatten(root *TreeNode) {
	res := new([]int)
	node := new(TreeNode)
	InBegin(root, res)

	for i := 0; i < len(*res); i++ {
		Insert(node, (*res)[i])
	}
	root = node.Right
}

//插入
func Insert(node *TreeNode, v int) {
	if node.Right != nil {
		Insert(node.Right, v)
	} else {
		node.Right = &TreeNode{v, nil, nil}
	}
}

//先序遍历
func InBegin(node *TreeNode, result *[]int) {
	*result = append(*result, node.Val)
	if node.Left != nil {
		InBegin(node.Left, result)
	}
	if node.Right != nil {
		InBegin(node.Right, result)
	}
}
