package main

import "fmt"

func main() {
	//-10, -3, 0, 5, 9
	head := new(ListNode)
	//head.Val = -10
	//head.Next = new(ListNode)
	//head.Next.Val = -3
	//head.Next.Next = new(ListNode)
	//head.Next.Next.Val = 0
	//head.Next.Next.Next = new(ListNode)
	//head.Next.Next.Next.Val = 5
	//head.Next.Next.Next.Next = new(ListNode)
	//head.Next.Next.Next.Next.Val = 9

	head.Val = -10
	head.Next = new(ListNode)
	head.Next.Val = -6
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = -3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = -2
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = -1
	head.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Val = 0
	head.Next.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Next.Val = 1
	head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Next.Next.Val = 2
	head.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val = 6
	//head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	//head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val = 10

	//fmt.Println(sortedListToBST(head))
	//root := sortedListToBST(head)
	//root := sortedListToBSTOfficial(head)
	root := sortedListToBSTOfficial1(head)
	res := []int{}
	fmt.Println(preorderTraversal1(root))
	fmt.Println(inBefore(root, res))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	root := new(TreeNode)

	var f func(arr []int, root *TreeNode)

	f = func(arr []int, root *TreeNode) {
		root.Val = arr[len(arr)/2]
		i, j := len(arr)/2-1, len(arr)/2+1
		if i >= 0 {
			root.Left = new(TreeNode)
			f(arr[:i+1], root.Left)
		}
		if j < len(arr) {
			root.Right = new(TreeNode)
			f(arr[j:], root.Right)
		}
	}

	f(arr, root)

	return root
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append([]int{root.Val}, preorderTraversal1(root.Left)...)
	rest = append(rest, preorderTraversal1(root.Right)...)
	return rest
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

func sortedListToBSTOfficial(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}

var globalHead *ListNode

func sortedListToBSTOfficial1(head *ListNode) *TreeNode {
	globalHead = head
	length := getLength(head)
	return buildTree1(0, length-1)
}

func getLength(head *ListNode) int {
	ret := 0
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildTree1(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &TreeNode{}
	root.Left = buildTree1(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree1(mid+1, right)
	return root
}
