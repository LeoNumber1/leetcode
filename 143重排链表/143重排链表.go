package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[1,2,3,4,5,6,7]"
	head := generateListNode(s)

	//reorderList(head)
	reorderListOfficial2(head)
}

//12 ms-76.73%	5.9 MB-18.83%
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	arr := []*ListNode{}
	node := head
	for node != nil {
		a := node
		arr = append(arr, a)
		node = node.Next
		a.Next = nil
	}

	pre := &ListNode{0, nil}
	n := len(arr)
	for i := 0; i < n; i++ {
		if i&1 != 0 { //奇数
			pre.Next = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			pre = pre.Next
		} else { //偶数
			pre.Next = arr[0]
			arr = arr[1:]
			pre = pre.Next
		}
	}

	printList(head)
}

//12 ms-76.73%	5.9 MB-18.83%
func reorderListOfficial1(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderListOfficial2(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
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

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d", l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}
}
