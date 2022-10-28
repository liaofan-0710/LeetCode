package main

import "fmt"

func minSubArrayLen(target int, nums []int) (res int) {
	if len(nums) < 2 && target == nums[0] {
		return 1
	}
	res = int(^uint(0) >> 1)
	i, j := 0, 1
	sum := nums[0]
	for i < j {
		if sum < target && j == len(nums) {
			break
		}
		if sum < target && j < len(nums) {
			sum += nums[j]
			j++
		}
		if sum >= target {
			res = min(res, j-i)
			sum -= nums[i]
			i++
		}
	}
	if res == int(^uint(0)>>1) {
		return 0
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
