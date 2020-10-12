package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	root := generateTree("[1,null,3,2]")
	root = generateTree("[236,104,701,null,227,null,911]")
	root = generateTree("[0,null,2236,1277,2776,519]")
	root = generateTree("[600,424,612,null,499,null,689]")
	//root = generateTree("[3,null,9,4]")

	//fmt.Println(getMinimumDifference(root))
	fmt.Println(getMinimumDifferenceMorris(root))
}

//16 ms-43.45%	6.3 MB-63.03%
func getMinimumDifference(root *TreeNode) int {
	var ans = math.MaxInt32
	var arr []int
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Left)
		arr = append(arr, node.Val)
		f(node.Right)
	}
	f(root)
	n := len(arr)
	for k, v := range arr {
		if k+1 < n {
			tmp := arr[k+1] - v
			if tmp < ans {
				ans = tmp
			}
		}
	}
	return ans
}

//8 ms-97.24%	6.5 MB-29.41%
func getMinimumDifferenceMorris(node *TreeNode) int {
	var ans = math.MaxInt32
	pre := -1
	for node != nil {
		if node.Left != nil {
			succ := getPre(node)
			if succ.Right == nil {
				succ.Right = node
				node = node.Left
			} else {
				if pre != -1 {
					tmp := node.Val - pre
					if tmp < ans {
						ans = tmp
					}
				}
				pre = node.Val
				//复原
				succ.Right = nil
				//遍历右子树
				node = node.Right
			}
		} else { //没有左子树
			if pre != -1 {
				tmp := node.Val - pre
				if tmp < ans {
					ans = tmp
				}
			}
			pre = node.Val
			node = node.Right
		}
	}

	return ans
}

//获取node节点的前继节点
func getPre(node *TreeNode) *TreeNode {
	succ := node.Left
	for succ.Right != nil && succ.Right != node {
		succ = succ.Right
	}
	return succ
}

//获取node节点的后继节点
func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

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
