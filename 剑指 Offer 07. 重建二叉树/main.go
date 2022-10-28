package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	// 拿前序数组首个数字，与中序对比，找到左右子树
	if len(preorder) == 0 {
		return nil
	}
	root := &tree.TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree.InorderTraversal(root))
}
