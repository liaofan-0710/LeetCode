package main

import (
	"fmt"
)

func countGood(nums []int, k int) (ans int64) {
	for l := 0; l < len(nums); l++ {
		countMap := make(map[int]int)
		for i := l; i < len(nums); i++ {
			countMap[nums[i]]++
		}
		count := 0
		for _, v := range countMap {
			if v == 2 {
				count++
			} else if v*2 >= k {
				ans++
			}
		}
		if count >= k {
			ans += int64(count-k) + 1
		}
	}
	return
}

func main() {
	fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
}
