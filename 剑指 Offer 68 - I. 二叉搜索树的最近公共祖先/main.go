package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

func lowestCommonAncestor(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode {
	if root.Left == p || root.Right == q {
		return root
	}
	lowestCommonAncestor(root.Left, p, q)
	lowestCommonAncestor(root.Right, p, q)
	return root
}

func main() {
	nums := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5, -1, -1, -1, -1}
	root := tree.SequenceCreation(nums)
	p := &tree.TreeNode{Val: 2}
	q := &tree.TreeNode{Val: 8}
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val)
}
