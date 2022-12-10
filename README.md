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

//二维数组按照第0位排序

sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

//先按照 cx<cy 排序，如果相等再按 x<y 排序
sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := onesCount(x), onesCount(y)
		return cx < cy || cx == cy && x < y
	})
```

力扣批量运行测试用例
```
type args struct {
    n           int
    connections string
}

tests := []struct {
    index  int
    args   args
    target int
}{
    {1, args{6, "[[0,1],[0,2],[0,3],[1,2],[1,3]]"}, 2},
    {2, args{6, "[[0,1],[0,2],[0,3],[1,2]]"}, -1},
    {3, args{5, "[[0,1],[0,2],[3,4],[2,3]]"}, 0},
    {4, args{4, "[[0,1],[0,2],[1,2]]"}, 1},
}

str2matrix := func(s string) (matrix [][]int) {
    arr := strings.Split(s, "],")
    for _, s2 := range arr {
        s2 = strings.TrimLeft(s2, "[")
        s2 = strings.TrimRight(s2, "]")
        arr1 := strings.Split(s2, ",")
        var temp []int
        for _, s3 := range arr1 {
            i, _ := strconv.Atoi(s3)
            temp = append(temp, i)
        }
        matrix = append(matrix, temp)
    }
    return
}

str2arr := func(s string) (res []int) {
    s = strings.TrimLeft(s, "[")
    s = strings.TrimRight(s, "]")
    arr := strings.Split(s, ",")
    for _, s2 := range arr {
        i, _ := strconv.Atoi(s2)
        res = append(res, i)
    }
    return
}

var errNum bool
for _, tt := range tests {
    result := makeConnected(tt.args.n, str2matrix(tt.args.connections))
    if tt.target != result {
        errNum = true
        fmt.Println("—————— err in index:", tt.index, "except:", tt.target, " get result:", result)
    }
}

if !errNum {
    fmt.Println("------- All tests are OK! -------")
}
```