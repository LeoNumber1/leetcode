package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 0
	root.Left = new(TreeNode)
	root.Left.Val = 0
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 2

	//fmt.Println(findMode1(root))
	//fmt.Println(findMode1(nil))
	//fmt.Println(findModeMorris(root))
	fmt.Println(findModeOfficial(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//20 ms	6.7 MB
func findMode(root *TreeNode) []int {
	m := map[int]int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		m[node.Val]++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	var ans []int
	var count int
	for k, v := range m {
		if v > count {
			count = v
			ans = []int{k}
		} else if v == count {
			ans = append(ans, k)
		}
	}
	return ans
}

//这个有问题，假如相同的值不是连在一起，就有问题了
func findMode1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		count int
		ans   []int
		dfs   func(node *TreeNode, num int)
	)
	dfs = func(node *TreeNode, num int) {
		if num > count {
			ans = []int{node.Val}
			count = num
		} else if num == count && node.Val != ans[len(ans)-1] {
			ans = append(ans, node.Val)
		}
		if node.Left != nil {
			if node.Val == node.Left.Val {
				dfs(node.Left, num+1)
			} else {
				dfs(node.Left, 1)
			}
		}
		if node.Right != nil {
			if node.Val == node.Right.Val {
				dfs(node.Right, num+1)
			} else {
				dfs(node.Right, 1)
			}
		}
	}
	dfs(root, 1)
	return ans
}

//12 ms-85.59%	6.2 MB-39.13%
func findModeMorris(root *TreeNode) []int {
	var base, count, maxCount int
	var ans []int
	for root != nil {
		if root.Left != nil {
			//左子树不为空，predecessor节点表示
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				//如果有右子树切没有设置过指向root，则继续往右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				//将predecessor的右指针指向root，这样遍历完root.left后，就能通过这个指针返回root
				predecessor.Right = root
				//遍历左子树
				root = root.Left
			} else {
				//predecessor的右指针已经指向了root，则表示左子树已经访问完了
				base, count, maxCount, ans = update(base, count, maxCount, root.Val, ans)

				//复原
				predecessor.Right = nil
				//遍历右子树
				root = root.Right
			}
		} else {
			//没有左子树，当前节点就是最小值了
			base, count, maxCount, ans = update(base, count, maxCount, root.Val, ans)
			root = root.Right
		}
	}

	return ans
}

func update(base, count, maxCount, val int, ans []int) (baseRet, countRet, maxCountRet int, ansRet []int) {
	if base != val {
		base = val
		count = 1
	} else {
		count++
	}
	if count > maxCount {
		ans = []int{val}
		maxCount = count
	} else if count == maxCount {
		ans = append(ans, val)
	}
	return base, count, maxCount, ans
}

//8 ms-99.15%	6.2 MB-39.13%
func findModeOfficial(root *TreeNode) (answer []int) {
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return
}
