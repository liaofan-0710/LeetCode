package main

import "fmt"

func main() {
	nums := []int{3,2,4}
	nums = twoSum(nums, 6)
	fmt.Println(nums)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{-1,-1}
}

//type ListNode struct {
//	Next *ListNode
//	Val  int
//}

//func main() {
//	//list := new(ListNode)
//	//fmt.Printf("%p\n", list)
//	//p := list
//	//p1 := list
//	//fmt.Printf("%p\n", p)
//	//fmt.Printf("%p\n", &p1)
//
//	//var num [1]int
//	//n := &num
//	//n1 := &num
//	//fmt.Printf("%p\n", &n)
//	//fmt.Printf("%p\n", &n1)
//	var str string
//	str = "aa"
//	fmt.Println(&str)
//	str = "aa" + "b"
//	fmt.Println(&str)
//	//fmt.Println(&p.Val)
//	//fmt.Println(&p1.Val)
//
//	//a := 0
//	//b := &a
//	//*b = 1
//	//fmt.Println(a)
//	//fmt.Printf("%p\n", &a)
//	//fmt.Printf("%p\n", a)
//	//fmt.Println(&a)
//	//fmt.Printf("%p\n", &b)
//}
