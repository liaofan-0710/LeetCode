package main

import "fmt"

func buildArray1(target []int, n int) []string {
	nums := []string{}
	for i, j := 1, 0; i <= n && j < len(target); i++ {
		nums = append(nums, "Push")
		if i != target[j] {
			nums = append(nums, "Pop")
		} else {
			j++
		}
	}
	return nums
}

func buildArray(target []int, n int) (ans []string) {
	prev := 0
	for i := 0; i < len(target); i++ {
		for j := 0; j < target[i]-prev-1; j++ {
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
		prev = target[i]
	}
	return
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 4))
}
