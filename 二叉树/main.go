package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			nums = append(nums, node.Val)
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			nums = append(nums, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums
}

// 后续遍历
func postorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			dfs(node.Right)
			nums = append(nums, node.Val)
		}
	}
	dfs(root)
	return nums
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

// 从中序后续创建二叉树
func iPBuildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
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

// 根据前序后续构建二叉树
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// 递归结束条件
	if len(preorder) == 0 {
		return nil
	}
	// 只有一个元素时，为叶子节点。这一步在889不可省略
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// 根节点
	root := &TreeNode{Val: preorder[0]}
	// 在后序序列中寻找和前序序列根节点下一个节点相等的位置
	for pos, node := range postorder {
		if node == preorder[1] {
			// 构建左子树
			root.Left = constructFromPrePost(preorder[1:pos+2], postorder[:pos+1])
			// 构建右子树
			root.Right = constructFromPrePost(preorder[pos+2:], postorder[pos+1:len(postorder)-1])
			break
		}
	}
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

func main() {
	// -1表示null
	nums := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := sequenceCreation(nums)
	result := levelOrder(root)
	fmt.Println(result)
}
