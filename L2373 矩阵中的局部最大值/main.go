package main

import "fmt"

// 暴力解法

// 非暴力解法，使用每三个为一组
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		res[i] = make([]int, n-2)
	}
	i, j := 0, 0
	for d := 1; d < n-1; d++ {
		for r := 1; r < n-1; r++ {
			if i < len(res) && j < len(res) {
				count1 := max(grid[d-1][r-1], grid[d-1][r], grid[d-1][r+1])
				count2 := max(grid[d][r-1], grid[d][r], grid[d][r+1])
				count3 := max(grid[d+1][r-1], grid[d+1][r], grid[d+1][r+1])
				res[j][i] = max(count1, count2, count3)
			}
			i++
		}
		j++
		i = 0
	}
	return res
}

func max(a, b, c int) int {
	count := 0
	if a > b {
		count = a
	} else {
		count = b
	}
	if count < c {
		count = c
	}
	return count
}

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
}
