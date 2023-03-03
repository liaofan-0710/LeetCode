package main

import (
	"fmt"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	i := 0
	for _, x := range nums[(len(nums)+1)/2:] {
		if nums[i]*2 <= x {
			i++
		}
	}
	return i * 2
}

func main() {
	fmt.Println(maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
}
