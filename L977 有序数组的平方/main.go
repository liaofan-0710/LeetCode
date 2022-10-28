package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-7,-3,2,3,11}
	result := sortedSquares(nums)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i], " ")
	}
}

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}