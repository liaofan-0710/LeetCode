package main

import "fmt"

func singleNumbers(nums []int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; !ok {
			count[nums[i]] = 1
		} else {
			count[nums[i]]++
		}
	}
	num := []int{0, 1}
	i := 0
	for k, v := range count {
		if v == 1 {
			num[i] = k
			i++
		}
	}
	return num
}

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}
