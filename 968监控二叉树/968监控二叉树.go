package main

import (
	"fmt"
	"math"
)

func main() {
	root := new(TreeNode)
	root.Left = new(TreeNode)
	root.Left.Left = new(TreeNode)
	root.Left.Right = new(TreeNode)
	root.Left.Right.Left = new(TreeNode)

	fmt.Println(minCameraCover(root))
	fmt.Println(minCameraCoverOfficial(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//有问题，不能算出特殊情况
/*
		0
	   / \
	  0   0
		   \
			0
			 \
			  0
			   \
				0
*/
func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nums := []int{}
	var total int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		total += n
		nums = append(nums, n)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	var ans, ans1 int
	n := len(nums)
	if n <= 2 {
		return nums[0]
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans += nums[i]
		}
		if (i-1)%3 == 0 {
			ans1 += nums[i]
		}
	}

	if n%3 == 0 { //n是3的倍数
		ans = min(ans, ans1)
	}
	return min(ans, total-ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const inf = math.MaxInt32 / 2 // 或 math.MaxInt64 / 2

func minCameraCoverOfficial(root *TreeNode) int {
	var dfs func(node *TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = min(a, min(la+rb, ra+lb))
		c = min(a, lb+rb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}
