package main

import "fmt"

func main() {
	//head := &ListNode{Val: 3}
	//head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 0}
	//head.Next.Next.Next = &ListNode{-4, head.Next}

	//head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 2, Next: head}

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{3, head}

	fmt.Println(detectCycle(head))
	fmt.Println(detectCycleOfficial(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//0 ms-100.00%	5.1 MB-13.01%
func detectCycle1(head *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	for head != nil {
		if _, has := m[head]; has {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

//双指针   4 ms-98.82%	3.7 MB-85.15%
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

//8 ms-76.29%	3.7 MB-85.15%
func detectCycleOfficial(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
