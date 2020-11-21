package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := "[4,2,1,3]"

	head := generateListNode(s)

	//printList(sortList(head))
	printList(sortListOfficial2(head))
}

//40 ms-28.59%	7.5 MB-14.47%
func sortList0(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	sort.Ints(arr)
	ans := new(ListNode)
	ans.Next = new(ListNode)
	head = ans.Next
	for k, v := range arr {
		head.Val = v
		if k != len(arr)-1 {
			head.Next = new(ListNode)
		}
		head = head.Next
	}
	return ans.Next
}

//插入排序，O（n2）
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := new(ListNode)
	list.Next = head
	curr, lastSorted := head.Next, head
	for curr != nil {
		if curr.Val >= lastSorted.Val {
			curr = curr.Next
			lastSorted = lastSorted.Next
			continue
		}
		prev := list
		for curr.Val > prev.Next.Val {
			prev = prev.Next
		}
		lastSorted.Next = curr.Next
		temp := prev.Next
		prev.Next = curr
		curr.Next = temp
		curr = lastSorted.Next
	}
	return list.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func mysort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(mysort(head, mid), mysort(mid, tail))
}

//归并排序，自顶向下，36 ms-41.16%	7.2 MB-27.93%
func sortListOfficial(head *ListNode) *ListNode {
	return mysort(head, nil)
}

func sortListOfficial2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
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
