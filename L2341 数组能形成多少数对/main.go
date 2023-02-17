package main

import "fmt"

func numberOfPairs(nums []int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	res := []int{0, 0}
	for _, v := range count {
		if v%2 == 0 {
			res[0] += v / 2
		} else {
			res[0] += v / 2
			res[1]++
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	fmt.Println(numberOfPairs(nums))
}
