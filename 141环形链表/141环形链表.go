package main

import "fmt"

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{-4, head.Next}

	fmt.Println(hasCycle(head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//进阶：快慢指针   8 ms-73.85% 	3.6 MB-85.16%
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}
	return false
}
