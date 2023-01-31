package main

import "fmt"

func checkXMatrix(grid [][]int) bool {
	n := len(grid) - 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if (i == j || i+j == n) && grid[i][j] == 0 {
				return false
			}
			if i != j && i+j != n && grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := [][]int{{5, 7, 0}, {0, 3, 1}, {0, 5, 0}}
	fmt.Println(checkXMatrix(nums))
}
