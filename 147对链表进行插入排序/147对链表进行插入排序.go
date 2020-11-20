package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[4,2,1,3]"
	s = "-1,5,3,4,0"

	head := generateListNode(s)
	//printList(insertionSortList(head))
	printList(insertionSortListOfficial(head))
}

//28 ms-20.34%	3.2 MB-81.36%
func insertionSortList(head *ListNode) *ListNode {
	node := head
	list := &ListNode{0, node}
	prevNode := list
out:
	for node != nil {
		prev := list
		a := list.Next
		next := node.Next
		for a != node {
			if node.Val < a.Val {
				//swap
				prevNode.Next = node.Next
				prev.Next = node
				node.Next = a
				node = next
				continue out
			}
			prev = a
			a = a.Next
		}
		node = next
		prevNode = prevNode.Next
	}
	return list.Next
}

//4 ms-96.90%	3.2 MB-39.78%
func insertionSortListOfficial(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateListNode(s string) (head *ListNode) {
	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")
	arr := strings.Split(s, ",")
	if len(arr) == 0 || arr[0] == "null" {
		return nil
	}
	head = new(ListNode)
	node := head
	for k, v := range arr {
		node.Val, _ = strconv.Atoi(v)
		if k != len(arr)-1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	return
}

//链表打印
func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}
}
