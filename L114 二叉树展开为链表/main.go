package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func flatten1(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}

func flatten2(root *TreeNode) {
	list := []*TreeNode{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		node = node.Right
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}
		left, right := curr.Left, curr.Right
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
		prev = curr
	}
}

func main() {
	nums := []int{1, 2, 5, 3, 4, -1, 6}
	//root := new(TreeNode)
	high, length, L := 1, 1, 1
	for i := 0; i < len(nums); i++ {
		length = int(math.Pow(2, float64(high-1)))
		list := nums[L : length+L]
		for j := 0; j < len(list); j++ {
			//root.Val
		}
		L += length
		high++
	}
}
