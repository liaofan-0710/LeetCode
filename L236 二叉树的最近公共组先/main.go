package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 层序创建二叉树
func sequenceCreation(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	// 先对数组进行分层分割
	sliNums := make([][]int, 0)
	count := 1
	for i := 0; i < len(nums); {
		cont := make([]int, 0)
		j := i
		for ; j < i+count; j++ {
			cont = append(cont, nums[j])
		}
		i = j
		sliNums = append(sliNums, cont)
		count *= 2
	}
	root := new(TreeNode)
	root.Val = sliNums[0][0]
	queue := []*TreeNode{root}
	count = 0
	for i := 1; i < len(sliNums); i++ {
		for j := 0; j < len(sliNums[i]); j += 2 {
			if sliNums[i][j] != -1 {
				queue[count].Left = &TreeNode{Val: sliNums[i][j]}
				queue = append(queue, queue[count].Left)
			} else {
				if queue[count] != nil {
					queue[count].Left = nil
					queue = append(queue, queue[count].Left)
				}
			}
			if sliNums[i][j+1] != -1 {
				queue[count].Right = &TreeNode{Val: sliNums[i][j+1]}
				queue = append(queue, queue[count].Right)
			} else {
				if queue[count] != nil {
					queue[count].Right = nil
					queue = append(queue, queue[count].Right)
				}
			}
			count++
		}
	}
	return root
}

func main() {
	nums := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4, -1, -1, -1, -1}
	root := sequenceCreation(nums)
	p := &TreeNode{
		Val: 5,
	}
	q := &TreeNode{
		Val: 1,
	}
	node := lowestCommonAncestor(root, p, q)
	fmt.Println(node.Val)

}
