package main

import "fmt"

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	//inorder = []int{3, 2, 1}
	//postorder = []int{3, 2, 1}

	//fmt.Println(show(buildTree(inorder, postorder)))
	fmt.Println(show(buildTree(inorder, postorder)))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//4 ms-96.47%	4.3 MB-5.02%
func buildTree1(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	var dfs func(inorder []int) *TreeNode
	dfs = func(inorder []int) *TreeNode {
		root := new(TreeNode)
		root.Val = postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		l, r := findIndex(inorder, root.Val)
		if len(r) > 0 {
			root.Right = dfs(r)
		}
		if len(l) > 0 {
			root.Left = dfs(l)
		}

		return root
	}

	return dfs(inorder)
}

func findIndex(inorder []int, val int) (left, right []int) {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			return inorder[:i], inorder[i+1:]
		}
	}
	return nil, nil
}

//0 ms-100.00%	4.2 MB-22.07%
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	i := n - 1
	for inorder[i] != postorder[n-1] {
		i--
	}
	return &TreeNode{
		postorder[n-1],
		buildTree(inorder[:i], postorder[:i]),
		buildTree(inorder[i+1:], postorder[i:n-1]),
	}
}

func InPostOrder(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	InPostOrder(node.Left, ans)
	InPostOrder(node.Right, ans)
	*ans = append(*ans, node.Val)
}

func show(node *TreeNode) []int {
	var ans []int
	if node == nil {
		return ans
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		this := queue[0]
		queue = queue[1:]
		ans = append(ans, this.Val)
		if this.Left != nil {
			queue = append(queue, this.Left)
		}
		if this.Right != nil {
			queue = append(queue, this.Right)
		}
	}
	return ans
}

func buildTreeOfficial(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}
