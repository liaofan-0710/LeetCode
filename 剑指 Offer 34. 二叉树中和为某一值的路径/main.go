package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

// 第一种：深搜每到最大深度判断
// 第二种：广搜

func pathSum(root *tree.TreeNode, target int) [][]int {
	result := make([][]int, 0)
	nums := make([]int, 0)
	var dfs func(node *tree.TreeNode, left int)
	dfs = func(node *tree.TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		nums = append(nums, node.Val)
		defer func() { nums = nums[:len(nums)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			result = append(result, append([]int(nil), nums...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, target)
	return result
}

func main() {
	nums := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1}
	root := tree.SequenceCreation(nums)
	fmt.Println(pathSum(root, 22))
}
