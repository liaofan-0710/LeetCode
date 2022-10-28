package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents1(head *ListNode, nums []int) int {
	ans := map[int]int{}
	res1 := [][]int{}
	res2 := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := ans[nums[i]]; ok {
			ans[nums[i]]++
		} else {
			ans[nums[i]] = 1
		}
	}
	for head != nil {
		if _, ok := ans[head.Val]; !ok {
			if len(res2) > 0 {
				res1 = append(res1, res2)
				res2 = []int{}
			}
			head = head.Next
		} else {
			if ans[head.Val] == 0 {
				if len(res2) > 0 {
					res1 = append(res1, res2)
					res2 = []int{}
				}
			} else {
				res2 = append(res2, head.Val)
				ans[head.Val]--
			}
			head = head.Next
		}
	}
	if len(res2) != 0 {
		res1 = append(res1, res2)
	}
	return len(res1)
}

func numComponents(head *ListNode, nums []int) (ans int) {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	for inSet := false; head != nil; head = head.Next {
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet {
			inSet = true
			ans++
		}
	}
	return
}

func main() {
	nums := []int{3, 4, 0, 2, 1}
	head := &ListNode{Val: nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		cur := &ListNode{Val: nums[i]}
		node.Next = cur
		node = node.Next
	}
	fmt.Println(numComponents(head, []int{4}))
}
