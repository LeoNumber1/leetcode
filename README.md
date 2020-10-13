# leetcode
leetcode算法题，go语言

力扣字符串生成二叉树代码
```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree(str string) (root *TreeNode) {
	s := strings.TrimLeft(str, "[")
	s = strings.TrimRight(s, "]")
	arr := strings.Split(s, ",")
	if len(arr) == 0 || arr[0] == "null" {
		return
	}
	root = new(TreeNode)
	root.Val, _ = strconv.Atoi(arr[0])
	arr = arr[1:]
	queue := []*TreeNode{root}
	for len(queue) > 0 && len(arr) > 0 {
		node := queue[0]
		queue = queue[1:]

		if arr[0] != "null" {
			node.Left = new(TreeNode)
			node.Left.Val, _ = strconv.Atoi(arr[0])
			queue = append(queue, node.Left)
		}
		arr = arr[1:]
		if len(arr) > 0 {
			if arr[0] != "null" {
				node.Right = new(TreeNode)
				node.Right.Val, _ = strconv.Atoi(arr[0])
				queue = append(queue, node.Right)
			}
			arr = arr[1:]
		}
	}
	return
}

//生成链表的代码

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
```