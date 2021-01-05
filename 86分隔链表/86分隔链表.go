package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[1,4,3,2,5,2]"
	x := 3

	head := generateListNode(s)

	//printList(partition(head, x))
	printList(partitionOfficial(head, x))
}

//0 ms-100.00%	2.4 MB-68.08%
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	var slow, fast *ListNode = dummy, head
	for fast != nil && fast.Val < x {
		slow = fast
		fast = fast.Next
	}
	if fast == nil {
		return head
	}
	var prev *ListNode = dummy
	for fast != nil {
		if fast.Val < x {
			temp := slow.Next
			slow.Next = fast
			prev.Next = fast.Next
			fast.Next = temp
			slow = slow.Next
			fast = prev.Next
		} else {
			prev = fast
			fast = fast.Next
		}
	}
	return dummy.Next
}

//0 ms	2.4 MB
func partitionOfficial(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
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
