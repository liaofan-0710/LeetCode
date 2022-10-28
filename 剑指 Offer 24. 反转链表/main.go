package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var per *ListNode
	node := head
	for node != nil {
		cur := node.Next
		node.Next = per
		per = node
		node = cur
	}
	return per
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
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}
