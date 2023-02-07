package main

import (
	"fmt"
	"sort"
)

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	res := 1
	for _, v := range coins {
		if v > res {
			break
		}
		res += v
	}
	return res
}

func main() {
	nums := []int{1, 3}
	fmt.Println(getMaximumConsecutive(nums))
}
