package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := tree.SequenceCreation(nums)
	fmt.Println(maxDepth(root))
}
