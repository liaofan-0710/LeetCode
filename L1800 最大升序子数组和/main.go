package main

import "fmt"

func maxAscendingSum1(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	result := 0
	count := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count += nums[i]
		} else {
			result = max(result, count)
			count = nums[i]
		}
	}
	if count > result {
		result = count
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxAscendingSum(nums []int) (ans int) {
	for i, n := 0, len(nums); i < n; {
		sum := nums[i]
		for i++; i < n && nums[i] > nums[i-1]; i++ {
			sum += nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

func main() {
	fmt.Println(maxAscendingSum([]int{3, 6, 10, 1, 8, 9, 9, 8, 9}))
}
