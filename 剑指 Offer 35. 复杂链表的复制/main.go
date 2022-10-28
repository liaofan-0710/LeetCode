package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}

func main() {
	nums := [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head := &Node{Val: nums[0][0]}
	// 第一步：复制原链表
	node := head
	for i := 1; i < len(nums); i++ {
		cur := &Node{Val: nums[i][0]}
		node.Next = cur
		node = node.Next
	}
	// 第二步：在原链表的基础上修改他的random
	node = head
	for i := 0; node != nil; i++ {
		if nums[i][1] != -1 {
			cur := &Node{Val: nums[i][1]}
			node.Random = cur
			node = node.Next
		} else {
			node.Random = nil
			node = node.Next
		}
	}
	res := copyRandomList(head)
	for res != nil {
		fmt.Println(res.Val)
		fmt.Println(res.Random)
		res = res.Next
	}
}
