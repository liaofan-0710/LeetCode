package main

import (
	"Project/tree"
	"fmt"
)

// 中顺遍历
func sortedArrayToBST(nums []int) *tree.TreeNode {
	l, r := 0, len(nums)-1
	return mediumOrder(l, r, nums)
}

func mediumOrder(l, r int, nums []int) *tree.TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	root := &tree.TreeNode{Val: nums[mid]}
	root.Left = mediumOrder(l, mid-1, nums)
	root.Right = mediumOrder(mid+1, r, nums)
	return root
}

// 内嵌函数递归
/*
func sortedArrayToBST(nums []int) *tree.TreeNode {
	var mediumOrder func(l, r int) *tree.TreeNode
	mediumOrder = func(l, r int) *tree.TreeNode {
		if l > r {
			return nil
		}
		mid := (l + r) / 2
		root := &tree.TreeNode{Val: nums[mid]}
		root.Left = mediumOrder(l, mid-1)
		root.Right = mediumOrder(mid+1, r)
		return root
	}

	return mediumOrder(0, len(nums)-1)
}
*/

func main() {
	fmt.Println(tree.InorderTraversal(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
}
