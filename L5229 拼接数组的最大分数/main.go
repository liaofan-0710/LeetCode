package main

import (
	"fmt"
)

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	//  只能执行一次操作
	// 将nums1
	n := len(nums1)
	getDiff := func(nums1, nums2 []int) int {
		// 将nums1 增加
		diffs := make([]int, n)
		sum := 0
		for i := 0; i < n; i++ {
			sum += nums1[i]
			diffs[i] = nums2[i] - nums1[i]
		}
		best := 0
		m := 0
		for i := 0; i < n; i++ {
			m += diffs[i]
			if m < 0 {
				m = 0
			}
			best = max(best, m)
		}
		return sum + best
	}
	return max(getDiff(nums1, nums2), getDiff(nums2, nums1))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{20, 40, 20, 70, 30}
	nums2 := []int{50, 20, 50, 40, 20}
	fmt.Println(maximumsSplicedArray(nums1, nums2))
}
