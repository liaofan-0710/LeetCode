package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if 2*sum+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
