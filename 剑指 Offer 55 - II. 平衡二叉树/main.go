package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	nums := []int{1, 2, 2, 3, 3, -1, -1, 4, 4, -1, -1, -1, -1, -1, -1}
	root := tree.SequenceCreation(nums)
	fmt.Println(isBalanced(root))
}
