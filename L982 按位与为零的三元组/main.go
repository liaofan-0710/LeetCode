package main

import "fmt"

func countTriplets(A []int) int {
	cnt := make(map[int]int)
	for _, a := range A {
		for _, b := range A {
			cnt[a&b]++
		}
	}
	ans := 0
	for k, v := range cnt {
		for _, a := range A {
			if (k & a) == 0 {
				ans += v
			}
		}
	}
	return ans
}

func main() {
	fmt.Println()
}
