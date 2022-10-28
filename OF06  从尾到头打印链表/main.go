package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	nums := make([]int, 0)
	for i := 0; head != nil; i++ {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

func main() {
	nums := []int{1, 3, 2}
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
	num := reversePrint(head)
	fmt.Println(num)
	//for head != nil {
	//	fmt.Println(head.val)
	//	head = head.Next
	//}
}
