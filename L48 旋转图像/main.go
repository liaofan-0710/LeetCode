package main

import (
	"fmt"
)

func rotate1(matrix [][]int) {
	// 首先拿一个切片存储第一行数字，之后进入循环后保存要覆盖掉的那一列数字，需要注意边界问题，列如：2*2的方阵，3*3的方阵，4*4的方阵中心是否需要继续递减覆盖
	contNum := make([]int, 0)
	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {
			contNum = append(contNum, matrix[i][j])
		}
		contNum = make([]int, 0)
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
