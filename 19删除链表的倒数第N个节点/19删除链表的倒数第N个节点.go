package main

import (
	"strconv"
	"strings"
)

func main() {
	s := "[1,2,3,4,5]"
	head := generateListNode(s)
	n := 5

	//removeNthFromEnd(head, n)
	//removeNthFromEndOfficial1(head, n)
	removeNthFromEndOfficial2(head, n)
}

//0 ms	2.2 MB
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	node := head
	var lenth int
	for node != nil {
		lenth++
		node = node.Next
	}
	node = head
	if n < lenth {
		for i := 0; i < lenth-n-1; i++ {
			node = node.Next
		}
		if node.Next != nil {
			node.Next = node.Next.Next
		}
	} else {
		head = head.Next
	}
	return head
}

//0 ms	2.4 MB
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lenth int
	m := map[int]*ListNode{}
	node := head
	for node != nil {
		m[lenth] = node
		lenth++
		node = node.Next
	}
	if lenth <= n {
		return head.Next
	}
	m[lenth-n-1].Next = m[lenth-n-1].Next.Next

	return head
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEndOfficial(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func removeNthFromEndOfficial1(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

//双指针	0 ms	2.2 MB
func removeNthFromEndOfficial2(head *ListNode, n int) *ListNode {
	first, second := &ListNode{0, head}, head
	dummy := first
	var count int
	for second != nil {
		if count >= n {
			first = first.Next
		}
		second = second.Next
		count++
	}
	first.Next = first.Next.Next
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
