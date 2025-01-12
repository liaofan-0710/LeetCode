package main

import "fmt"

// 双指针
func trap(height []int) (ans int) {
	leftMax, rightMax := 0, 0
	for l, r := 0, len(height)-1; l < r; {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		if height[l] < height[r] {
			ans += leftMax - height[l]
			l++
		} else {
			ans += rightMax - height[r]
			r--
		}

	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
