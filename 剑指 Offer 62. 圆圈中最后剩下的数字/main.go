package main

import "fmt"

func lastRemaining(n int, m int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	count := m
	for len(nums) > 1 {
		nums = append(nums[:count-1], nums[count:]...)
		count = len(nums[count-1:])
		count = m - count
		count %= len(nums)
	}
	return nums[0]
}

func main() {
	fmt.Println(lastRemaining(5, 3))
}
