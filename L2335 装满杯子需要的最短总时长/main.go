package main

import (
	"fmt"
	"sort"
)

func fillCups(amount []int) int {
	res := 0
	for amount[0] != 0 || amount[1] != 0 || amount[2] != 0 {
		sort.Ints(amount)
		if amount[2] > 0 {
			amount[2]--
		}
		if amount[1] > 0 {
			amount[1]--
		}
		res++
	}
	return res
}

func main() {
	fmt.Println(fillCups([]int{1, 4, 2}))
}
