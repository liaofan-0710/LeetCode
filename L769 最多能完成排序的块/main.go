package main

import "fmt"

func maxChunksToSorted1(arr []int) int {
	count, sum := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > sum {
			sum = arr[i]
		}
		if sum == i {
			count++
		}
	}
	return count
}

func maxChunksToSorted(arr []int) int {
	n, ans := len(arr), 0
	for i, j, min1, max1 := 0, 0, n, -1; i < n; i++ {
		min1 = min(min1, arr[i])
		max1 = max(max1, arr[i])
		if j == min1 && i == max1 {
			ans++
			j = i + 1
			min1 = n
			max1 = -1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}
