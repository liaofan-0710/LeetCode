package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据前序后续构建二叉树
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{postorder[0], nil, nil}
	}
	root := new(TreeNode)
	i := 0
	for i = 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			break
		}
	}
	lenLeft := i + 1
	root.Val = preorder[0]
	root.Left = constructFromPrePost(preorder[1:1+lenLeft], postorder[:lenLeft])
	root.Right = constructFromPrePost(preorder[1+lenLeft:], postorder[lenLeft:len(postorder)-1])
	return root
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
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	root := constructFromPrePost(preorder, postorder)
	post := postorderTraversal(root)
	fmt.Print("\n", post)
}
