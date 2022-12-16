package main

import "fmt"

func minElements(nums []int, limit, goal int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return (abs(sum-goal) + limit - 1) / limit
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minElements([]int{0, 0, 1, -1, 1}, 10, 0))
}
