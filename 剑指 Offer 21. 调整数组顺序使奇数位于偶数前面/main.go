package main

import "fmt"

func exchange1(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]%2 != 0 {
			i++
			continue
		}
		if nums[j]%2 == 0 {
			j--
			continue
		}
		if nums[j]%2 != 0 && nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func exchange(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		for i < j && (nums[i]&1) == 1 {
			i++
		}
		for i < j && (nums[j]&1) == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}
