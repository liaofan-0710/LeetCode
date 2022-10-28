package main

import (
	"fmt"
)

func main() {
	nums := []int{-1,0}
	target := -1
	result := twoSum(nums,target)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i] , " ")
	}
}

/*
思想：先找比 target 小的数字，最终用双指针相加
*/
func twoSum(numbers []int, target int) []int {
	// 1. 用二分搜索截取比 target小的数组numbers
	//n := len(numbers)
	//left, right := 0, n - 1
	//ans := n
	//for left <= right {
	//	mid := (right - left) >> 1 + left
	//	if numbers[mid] < 0 {
	//		if target*-1 <= numbers[mid]*-1 {
	//			ans = mid
	//			right = mid - 1
	//		} else {
	//			left = mid + 1
	//		}
	//	}
	//	if target <= numbers[mid] {
	//		ans = mid
	//		right = mid - 1
	//	} else {
	//		left = mid + 1
	//	}
	//}
	//numbers = numbers[0:ans]
	// 2. 用双指针暴力查找
	for i := 0; i < len(numbers); i++ {
		for j := i+1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				numbers = []int{i+1,j+1}
			}
		}
	}
	return numbers
}