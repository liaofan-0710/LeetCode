package main

import "fmt"

func missingNumber1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return nums[len(nums)-1] + 1
}

func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

func main() {
	fmt.Println(missingNumber([]int{0}))
}
