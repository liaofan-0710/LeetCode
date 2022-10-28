package main

import "fmt"

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for nums[i]+nums[j] != target {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		}
	}
	return []int{nums[i], nums[j]}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
