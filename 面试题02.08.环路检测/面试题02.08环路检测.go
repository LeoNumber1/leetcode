package main

import "fmt"

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{-4, head.Next}

	fmt.Println(detectCycle(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//此题和142环形链表II是同一道题

//双指针，	8 ms	3.7 MB
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	if head.Next.Next == head { //两个节点成环
		return head
	}
	var slow, fast = head.Next, head.Next.Next
	var ptr = head
	if fast == nil {
		return nil
	}
	for slow != fast {
		slow = slow.Next
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
	}
	for slow != ptr {
		slow = slow.Next
		ptr = ptr.Next
	}
	return ptr
}
