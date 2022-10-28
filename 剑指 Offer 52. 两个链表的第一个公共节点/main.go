package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

func main() {
	nums1 := []int{4, 1, 8, 4, 5}
	nums2 := []int{5, 0, 1, 8, 4, 5}
	head1 := &ListNode{Val: nums1[0]}
	head2 := &ListNode{Val: nums2[0]}
	node1, node2 := head1, head2
	for i := 1; i < len(nums1); i++ {
		cur := &ListNode{Val: nums1[i]}
		node1.Next = cur
		node1 = node1.Next
	}
	for i := 1; i < len(nums2); i++ {
		cur := &ListNode{Val: nums2[i]}
		node2.Next = cur
		node2 = node2.Next
	}
	fmt.Println(getIntersectionNode(head1, head2))
}
