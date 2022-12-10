package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[1,2,3,4,5,6,7,8,9,10]"
	head := generateListNode(s)
	k := 10

	printList(rotateRight(head, k))
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var n int
	last, temp := new(ListNode), head

	for temp != nil {
		n++
		last = temp
		temp = temp.Next
	}
	k %= n
	if k == 0 {
		return head
	}
	var prev = new(ListNode)
	prev.Next = head
	for i := 0; i < n-k; i++ {
		prev = prev.Next
	}

	dummy := new(ListNode)
	dummy.Next = prev.Next
	prev.Next = nil
	last.Next = head

	return dummy.Next
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
