package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println(result)
}

func maxSubArray1(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, j := range nums {
		if sum > 0 {
			sum += j
		} else {
			sum = j
		}
		if ans < sum {
			ans = sum
		}
	}
	return ans
}

func maxSubArray(nums []int) int {
	max := nums[0]
	// 模拟窗口滑动
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}