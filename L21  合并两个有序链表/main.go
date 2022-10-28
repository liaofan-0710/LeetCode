package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	l1 := []int{1,2,4}
	l2 := []int{1,3,4}
	head1 := &ListNode{
		Val: l1[0],
	}
	head2 := &ListNode{
		Val: l2[0],
	}
	node1 := head1
	node2 := head2
	for i := 1; i < len(l1); i++ {
		node := new(ListNode)
		node.Val = l1[i]
		node1.Next = node
		node1 = node1.Next
	}
	for i := 1; i < len(l2); i++ {
		node := new(ListNode)
		node.Val = l2[i]
		node2.Next = node
		node2 = node2.Next
	}
	node := head1
	node = mergeTwoLists(head1, head2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	pnode := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node := &ListNode{Val: l1.Val}
			pnode.Next = node
			l1 = l1.Next
		}else{
			node := &ListNode{Val: l2.Val}
			pnode.Next = node
			l2 = l2.Next
		}
		pnode = pnode.Next
	}
	if l1 != nil {
		pnode.Next = l1
	}
	if l2 != nil{
		pnode.Next = l2
	}
	return head.Next
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head,next,cur *ListNode = nil, nil, nil
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur = l2
			l2 = l2.Next
		} else if l2 == nil {
			cur = l1
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				cur = l1
				l1 = l1.Next
			} else {
				cur = l2
				l2 = l2.Next
			}
		}
		if head == nil {
			head = cur
			next = head
		} else {
			next.Next = cur
			next = cur
		}
	}
	return head
}