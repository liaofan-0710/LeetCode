package main

import (
	"fmt"
	"sort"
)

// 双指针
func search1(nums []int, target int) int {
	result := 0
	for i, j := 0, len(nums)-1; i <= j; {
		if i != j && nums[i] == target {
			result++
		}
		if nums[j] == target {
			result++
		}
		i++
		j--
	}
	return result
}

// 一头找尾，一头找开始，进行减法
func search(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}

func main() {
	fmt.Println(search([]int{1, 3, 4}, 4))
}
