package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

func levelOrder(root *tree.TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*tree.TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*tree.TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := tree.SequenceCreation(nums)
	fmt.Println(tree.PreorderTraversal(root))
	fmt.Println(levelOrder(root))
}
