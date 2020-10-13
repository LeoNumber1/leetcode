package main

import "fmt"

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4

	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = new(ListNode)
	l2.Next.Val = 3
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	l2.Next.Next.Next = new(ListNode)
	l2.Next.Next.Next.Val = 5

	//printList(mergeTwoLists2(l1, l2))
	printList(mergeTwoLists4(l1, l2))
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d", l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	ln := new(ListNode)
	ln.Next = new(ListNode)

	sn := ln.Next
	//sn = new(ListNode)
	for l1 != nil || l2 != nil {
		//sn = new(ListNode)
		if l1 == nil {
			sn.Val = l2.Val
			l2 = l2.Next
			if l2 != nil {
				sn.Next = new(ListNode)
				sn = sn.Next
			}
			continue
		}
		if l2 == nil {
			sn.Val = l1.Val
			l1 = l1.Next
			if l1 != nil {
				sn.Next = new(ListNode)
				sn = sn.Next
			}
			continue
		}

		if l1.Val > l2.Val {
			sn.Val = l2.Val
			l2 = l2.Next
		} else {
			sn.Val = l1.Val
			l1 = l1.Next
		}
		if l1 != nil || l2 != nil {
			sn.Next = new(ListNode)
			sn = sn.Next
		}
	}
	return ln.Next
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	ln := new(ListNode)

	sn := ln
	for l1 != nil || l2 != nil {

		if l1 == nil {
			sn.Next = l2
			break
		}
		if l2 == nil {
			sn.Next = l1
			break
		}

		sn.Next = new(ListNode)
		sn = sn.Next

		if l1.Val > l2.Val {
			sn.Val = l2.Val
			l2 = l2.Next
		} else {
			sn.Val = l1.Val
			l1 = l1.Next
		}
	}
	return ln.Next
}

//func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil && l2 == nil {
//		return nil
//	}
//	ln := l1
//	for l1 != nil {
//		for l2 != nil {
//			if l1.Val < l2.Val {
//
//			}
//		}
//	}
//}

func mergeTwoLists4(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return result.Next
}
