package main

import (
	tree "LeetCode/Tree"
	"fmt"
)

// 广度遍历
func levelOrder(root *tree.TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*tree.TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := tree.SequenceCreation(nums)
	fmt.Println(levelOrder(root))
}
