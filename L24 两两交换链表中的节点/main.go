package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	nums := []int{1}
	head := &ListNode{
		Val : nums[0],
	}
	node := head
	for i := 1; i < len(nums); i++ {
		cur := new(ListNode)
		cur.Val = nums[i]
		node.Next = cur
		node = node.Next
	}
	result := swapPairs(head)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := &ListNode{}

	cur = head.Next
	head.Next = cur.Next
	cur.Next = head

	cur.Next.Next = swapPairs(cur.Next.Next)
	return cur
}