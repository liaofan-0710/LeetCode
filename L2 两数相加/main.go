package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	cur := result
	digit1, digit2 := 0, 0
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 1 {
			node := &ListNode{Val: 1}
			cur.Next = node
			cur = cur.Next
			carry = 0
		}
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			digit1 = l1.Val
			l1 = l1.Next
		} else {
			digit1 = 0
		}
		if l2 != nil {
			digit2 = l2.Val
			l2 = l2.Next
		} else {
			digit2 = 0
		}
		node := &ListNode{Val: (digit1 + digit2 + carry) % 10}
		cur.Next = node
		cur = cur.Next
		if digit1+digit2+carry > 9 {
			carry = 1
		} else {
			carry = 0
		}
	}
	return result.Next
}

func main() {
	nums1 := []int{9, 9, 9, 9, 9, 9, 9}
	nums2 := []int{9, 9, 9, 9}
	head1 := &ListNode{Val: nums1[0]}
	head2 := &ListNode{Val: nums2[0]}
	cur1, cur2 := head1, head2
	for i := 1; i < len(nums1); i++ {
		node := &ListNode{Val: nums1[i]}
		cur1.Next = node
		cur1 = cur1.Next
	}
	for i := 1; i < len(nums2); i++ {
		node := &ListNode{Val: nums2[i]}
		cur2.Next = node
		cur2 = cur2.Next
	}
	result := addTwoNumbers(head1, head2)
	for result != nil {
		fmt.Printf(" %d", result.Val)
		result = result.Next
	}
	//for head1 != nil {
	//	fmt.Print(head1.Val)
	//	head1 = head1.Next
	//}
	//fmt.Println()
	//for head2 != nil {
	//	fmt.Print(head2.Val)
	//	head2 = head2.Next
	//}
}
