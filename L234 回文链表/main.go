package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	hash1, hash2, h := 0, 0, 1
	seed := 4
	p := head
	for p != nil {
		hash1 = hash1*seed + p.Val
		hash2 = hash2 + h*p.Val
		h *= seed
		p = p.Next
	}
	return hash1 == hash2
}

func main() {
	nums := []int{1, 2, 2, 1}
	head := &ListNode{Val: nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		cur := new(ListNode)
		cur.Val = nums[i]
		node.Next = cur
		node = node.Next
	}
	fmt.Println(isPalindrome(head))
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
