package main

import "fmt"

func main() {
	l1, l2 := new(ListNode), new(ListNode)

	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 6

	printList(addTwoNumbers(l1, l2))
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//12 ms	4.9 MB
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res = new(ListNode)
	r := res
	var addRes int

	for l1 != nil || l2 != nil || addRes > 0 {
		r.Next = new(ListNode)
		r = r.Next
		if l1 != nil {
			addRes += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			addRes += l2.Val
			l2 = l2.Next
		}

		r.Val = addRes % 10
		addRes /= 10
	}

	return res.Next
}

//28 ms	5.2 MB
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{Val: sum, Next: nextNode}
	} else {
		tempNode := &ListNode{
			Val:  1,
			Next: nil,
		}
		return &ListNode{
			Val:  sum - 10,
			Next: addTwoNumbers(nextNode, tempNode),
		}
	}
}
