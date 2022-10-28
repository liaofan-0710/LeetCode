package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	fmt.Println(root)
	return []int{}
}

func preOrder(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Print(root.Val, " ")
	preOrder(root.Left)
	preOrder(root.Right)
}

func start(root *TreeNode, nums []int, index int) {
	if nums[index] == -1 {
		root = nil
		return
	}
	//root = new(TreeNode)
	root.Val = nums[index]
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	start(root.Left, nums, index+1)
	start(root.Right, nums, index+2)
	preOrder(root)
	fmt.Println()
}

func main() {
	nums := []int{1, -1, 2, -1, 3, -1, -1}
	head := new(TreeNode)
	start(head, nums, 0)
	preOrder(head)
	//nums = inorderTraversal(head)
}
