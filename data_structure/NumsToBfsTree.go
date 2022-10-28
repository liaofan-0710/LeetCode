package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 前序遍历
func PreOrder(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	res = append(res, PreOrder(root.Left)...)
	res = append(res, PreOrder(root.Right)...)
	return res
}

// MidOrder 中序遍历
func MidOrder(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, MidOrder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, MidOrder(root.Right)...)
	return res
}

// PostOrder 后序遍历
func PostOrder(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, PostOrder(root.Left)...)
	res = append(res, PostOrder(root.Right)...)
	res = append(res, root.Val)
	return res
}

// LevelOrder 层次遍历
// 先将根节点放入队列
// 记录当前节点的Val，出队列，如果左子树有孩子，入队，如果右子树有孩子，入队
// 重复操作2，直到队列为空
func levelOrder(root *TreeNode) [][]int {
	// 定义返回二维数组
	ret := [][]int{}
	if root == nil {
		return ret
	}
	// 记录每一层的节点
	temp := []int{}
	// 根节点入队
	queue := []*TreeNode{root}
	// 队列不为空
	for len(queue) > 0 {
		// 后续有入队操作，保存当前len(queue)
		length := len(queue)
		// 更新
		temp = []int{}
		for i := 0; i < length; i++ {
			// 记录当前节点
			node := queue[0]
			// 出队
			queue = queue[1:]
			// 如果有左孩子，入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 如果有右孩子，入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			// 记录当前节点Val
			temp = append(temp, node.Val)
		}
		ret = append(ret, temp)
	}
	return ret
}

// 前序找根节点，中序找左右分界线
func pIBuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = pIBuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = pIBuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := pIBuildTree(preorder, inorder)
	nums := levelOrder(root)
	fmt.Println(nums)
}
