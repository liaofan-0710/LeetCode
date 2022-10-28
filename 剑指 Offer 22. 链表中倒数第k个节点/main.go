package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	for ; length > k; length-- {
		head = head.Next
	}
	return head
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
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
	res := getKthFromEnd(head, 2)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}
