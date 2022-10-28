package main

import (
	"fmt"
	"math"
)

func main()  {
	nums := []int{1,8,6,2,5,4,8,3,7}
	result := maxArea(nums)
	fmt.Println(result)
}

func maxArea(height []int) int {
	var max int = 0
	for i, j := 0, len(height)-1; i < j; {
		minHeight := 0
		if height[i] < height[j] {
			i ++
			minHeight = height[i]
		}else{
			j --
			minHeight = height[j]
		}
		area := (j - i + 1) * minHeight
		max = int(math.max(float64(area), float64(max)))
	}
}