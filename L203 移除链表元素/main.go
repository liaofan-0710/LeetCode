package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		// 中间部分的一位符合条件
		if cur.Next.Val == val {
			node := cur.Next.Next
			cur.Next = node
		} else {
			cur = cur.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = cur.Next
	}
	result := removeElements(head, 6)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
