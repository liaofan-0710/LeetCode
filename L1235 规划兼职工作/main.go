package main

import (
	"fmt"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	length := len(profit)
	nums := make([][3]int, length)
	for i := 0; i < length; i++ {
		nums[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		k := sort.Search(i, func(j int) bool { return nums[j][1] > nums[i-1][0] })
		dp[i] = max(dp[i-1], dp[k]+nums[i-1][2])
	}
	return dp[length]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
}
