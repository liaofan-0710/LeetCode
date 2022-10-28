package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	nums := []int{1,2,3,4,5,6}
	head := &ListNode{
		Val: nums[0],
	}
	node := head
	for i := 1; i < len(nums); i++ {
		tmp := new(ListNode)
		tmp.Val = nums[i]
		node.Next = tmp
		node = node.Next
	}
	pnode := middleNode(head)
	for pnode != nil {
		fmt.Println(pnode.Val)
		pnode = pnode.Next
	}
}

func middleNode(head *ListNode) *ListNode {
	tail := head
	for tail != nil && tail.Next != nil {
		head = head.Next
		tail = tail.Next
		tail = tail.Next
	}
	return head
}