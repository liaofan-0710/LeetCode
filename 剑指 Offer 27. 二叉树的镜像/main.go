package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = mirrorTree(root.Left), mirrorTree(root.Right)
	return root
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

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
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
	nums := []int{4, 2, 7, 1, 3, 6, 9}
	root := sequenceCreation(nums)
	res := mirrorTree(root)
	fmt.Println(levelOrder(res))
}
