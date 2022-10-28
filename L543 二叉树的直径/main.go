package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	high := &m
	dept(root, high)
	return *high
}
func dept(node *TreeNode, high *int) int {
	if node == nil {
		return 0
	}
	l := dept(node.Left, high)
	r := dept(node.Right, high)
	*high = int(math.Max(float64(*high), float64(l+r)))
	return int(math.Max(float64(l), float64(r))) + 1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	root := &TreeNode{Val: nums[0]}
	cur := root
	count, index, high, rootHigh := 2, 1, math.Sqrt(float64(len(nums)))+1, 2
	for rootHigh < int(high) {
		for i := index; i < count+index; i++ {
			if i <= count+index/2 { // +left
				node := new(TreeNode)
				node.Left.Val = nums[i]
				cur.Left = node.Left
				cur = cur.Left
			} else { // +right

			}
		}
	}
}
