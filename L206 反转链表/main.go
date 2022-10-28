package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := &ListNode{
		Val: nums[0],
	}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := new(ListNode)
		node.Val = nums[i]
		cur.Next = node
		cur = cur.Next
	}
	result := reverseList(head)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(f(0))
}

func f(i int) int {
	if i > 100 {
		return 0
	}
	return f(i+1) + i
}

func reverseList1(head *ListNode) *ListNode {
	node := head
	var nums []int
	result := &ListNode{}
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	node = result
	for i := len(nums) - 1; i >= 0; i-- {
		cur := new(ListNode)
		cur.Val = nums[i]
		node.Next = cur
		node = node.Next
	}
	return result.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
