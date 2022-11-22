package main

import (
	"fmt"
	"math"
)

func largestAltitude(gain []int) int {
	max := math.MinInt
	count := 0
	for i := 0; i < len(gain); i++ {
		gain[i] = gain[i] + count
		count = gain[i]
		if max < count {
			max = count
		}
	}
	if max < 0 {
		max = 0
	}
	return max
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
}
