package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "[1,2,3,4,5,6,7,8,9]"
	//s = "[2,1,3,5,6,4,7]"
	//s = "1,2,3"

	//printList(oddEvenList(generateListNode(s)))
	printList(oddEvenListOfficial(generateListNode(s)))
}

//执行耗时:4 ms,击败了87.97% 的Go用户
//内存消耗:3.2 MB,击败了38.55% 的Go用户
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var slow, fast = head, head.Next
	for fast != nil && fast.Next != nil {
		temp := fast.Next.Next
		slowNext := slow.Next
		slow.Next = fast.Next
		slow.Next.Next = slowNext
		fast.Next = temp
		slow = slow.Next
		fast = fast.Next
	}

	return head
}

//执行耗时:4 ms,击败了87.97% 的Go用户
//内存消耗:3.2 MB,击败了38.55% 的Go用户
func oddEvenListOfficial(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
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
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}
}
