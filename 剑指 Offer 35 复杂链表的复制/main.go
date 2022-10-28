package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//func copyRandomList(head *Node) *Node {
//	head = copy(head)
//	return head
//}
//
//func main() {
//	nums := [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
//	// 需要记录一个指针
//	head := Node{
//		Val:    nums[0][0],
//		Random: nil,
//	}
//	node := head
//	for i := 0; i < len(nums); i++ {
//		cur := new(Node)
//		cur.Val = nums[i]
//	}
//	copyRandomList(&head)
//}
