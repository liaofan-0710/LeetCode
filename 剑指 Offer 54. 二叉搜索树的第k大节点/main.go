package main

import (
	tree "LeetCode/Tree"
	"fmt"
	"sort"
)

// 第一种：将所有节点的值存储在数组中进行排序，根据下标返回
func kthLargest(root *tree.TreeNode, k int) int {
	nums := make([]int, 0)
	dfs := func(root *tree.TreeNode) {}
	dfs = func(root *tree.TreeNode) {
		if root != nil {
			nums = append(nums, root.Val)
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	sort.Sort(sort.IntSlice(nums))
	return nums[len(nums)-k]
}

func main() {
	nums := []int{3, 1, 4, -1, 2, -1, -1}
	root := tree.SequenceCreation(nums)
	fmt.Println(kthLargest(root, 1))
}
