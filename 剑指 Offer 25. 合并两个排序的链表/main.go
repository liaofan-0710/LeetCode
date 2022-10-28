package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	node := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur := &ListNode{Val: l1.Val}
			node.Next = cur
			node = node.Next
			l1 = l1.Next
		} else {
			cur := &ListNode{Val: l2.Val}
			node.Next = cur
			node = node.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		node.Next = l1
	} else if l2 != nil {
		node.Next = l2
	}
	return head.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
	return l1
}

func main() {
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4}
	head1 := &ListNode{Val: nums1[0]}
	head2 := &ListNode{Val: nums2[0]}
	node1 := head1
	node2 := head2
	for i := 1; i < len(nums1); i++ {
		cur1 := &ListNode{Val: nums1[i]}
		node1.Next = cur1
		node1 = node1.Next
		cur2 := &ListNode{Val: nums2[i]}
		node2.Next = cur2
		node2 = node2.Next
	}
	head := mergeTwoLists(head1, head2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
