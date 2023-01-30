package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween1(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	node1 := list1
	a1, b2 := 0, 0
	for node1 != nil {
		if a1+1 == a {
			cur1 := node1
			b2 = a1
			for b2 != b {
				node1 = node1.Next
				b2++
			}
			cur1.Next = list2
			cur2 := node1.Next
			for cur1.Next != nil {
				cur1 = cur1.Next
			}
			cur1.Next = cur2
			break
		} else {
			node1 = node1.Next
			a1++
		}
	}
	return list1
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	preA := list1
	for i := 0; i < a-1; i++ {
		preA = preA.Next
	}
	preB := preA
	for i := 0; i < b-a+2; i++ {
		preB = preB.Next
	}
	preA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = preB
	return list1
}

func main() {
	list1 := &ListNode{}
	list2 := &ListNode{}
	nums1 := []int{0, 3, 2, 1, 4, 5}
	nums2 := []int{1000000, 1000001, 1000002}
	cur1, cur2 := list1, list2
	for i := 0; i < len(nums1); i++ {
		node := &ListNode{Val: nums1[i]}
		cur1.Next = node
		cur1 = cur1.Next
	}
	for i := 0; i < len(nums2); i++ {
		node := &ListNode{Val: nums2[i]}
		cur2.Next = node
		cur2 = cur2.Next
	}
	result := mergeInBetween(list1.Next, 3, 4, list2.Next)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
