package main

import "fmt"

func main(){
	nums := []int{4,5,6,7,0,1,2}
	result := search(nums, 0)
	fmt.Println(result)

}

// 暴力扫描
func search1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// 二分搜索
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) ||
			(nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}