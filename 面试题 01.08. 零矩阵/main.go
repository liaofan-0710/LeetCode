package main

import (
	"fmt"
)

func setZeroes1(matrix [][]int) {
	nums := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				nums = append(nums, []int{i, j})
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(matrix[nums[i][0]]); j++ {
			matrix[nums[i][0]][j] = 0
		}
		for j := 0; j < len(matrix); j++ {
			matrix[j][nums[i][1]] = 0
		}
	}
}

func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

func setZeroes3(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

func main() {
	nums := [][]int{
		{0, 0, 0, 5},
		{4, 3, 1, 4},
		{0, 1, 1, 4},
		{1, 2, 1, 3},
		{0, 0, 1, 1},
	}
	setZeroes(nums)
	fmt.Println(nums)
}
