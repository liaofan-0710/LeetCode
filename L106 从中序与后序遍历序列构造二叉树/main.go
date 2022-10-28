package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序后续创建二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
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

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			nums = append(nums, node.Val)
			fmt.Print(node.Val, " ")
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums
}

// 后续遍历
func postorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			dfs(node.Right)
			nums = append(nums, node.Val)
			fmt.Print(node.Val, " ")
		}
	}
	dfs(root)
	return nums
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	inor := postorderTraversal(root)
	fmt.Print("\n", inor)
}
