package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

func levelOrder(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := make([]int, 0)
	q := []*tree.TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*tree.TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			nums = append(nums, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return nums
}

func main() {
	nums := []int{}
	root := tree.SequenceCreation(nums)
	fmt.Println(levelOrder(root))
}
