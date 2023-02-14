package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		cur := &ListNode{Val: nums[i]}
		node.Next = cur
		node = node.Next
	}
	res := reverseList(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
