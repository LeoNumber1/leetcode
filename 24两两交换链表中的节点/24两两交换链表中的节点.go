package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	head := generateListNode("[1,2,3,4,5]")

	//printList(swapPairs(head))
	//printList(swapPairsOfficial(head))
	printList(swapPairsOfficial1(head))
}

//0 ms-100.00%	2.1 MB-29.15%
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	node := head
	head = head.Next
	for node != nil && node.Next != nil {
		if pre != nil {
			pre.Next = node.Next
		}
		a := node.Next
		b := node.Next.Next
		a.Next = node
		pre = node
		node.Next = b
		node = b
	}

	return head
}

//	0 ms-100.00%	2.1 MB-46.41%
func swapPairsOfficial(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairsOfficial(newHead.Next)
	newHead.Next = head
	return newHead
}

//	0 ms-100.00%	2.1 MB-46.41%
func swapPairsOfficial1(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
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
