package main

type ListNode struct {
    Val int
  	Next *ListNode
}

func main() {
	nums := []int{1,2,3,4,5}
	head := &ListNode{
		Val: nums[0],
	}
	pnode := head
	for i := 1; i < len(nums); i++ {
		node := new(ListNode)
		node.Val = nums[i]
		pnode.Next = node
		pnode = pnode.Next
	}
	removeNthFromEnd(head, 2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenght := 0
	node := head
	for node != nil {
		lenght++
		node = node.Next
	}
	lenght -= n
	if lenght == 0 {
		return head.Next
	}
	node = head
	for i := 1; i < lenght; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}

// 双指针
//func removeNthFromEnd1(head *ListNode, n int) *ListNode {
//
//}