package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	nums := make([]int, len(matrix)*len(matrix[0]))
	index := 0
	for index < len(nums) {
		for i := 0; i < len(matrix[0]); i++ {
			nums[index] = matrix[0][i]
			index++
		}
		if index >= len(nums) {
			break
		}
		matrix = matrix[1:][:]
		for i := 0; i < len(matrix); i++ {
			nums[index] = matrix[i][len(matrix[i])-1]
			matrix[i] = matrix[i][:len(matrix[i])-1]
			index++
		}
		if index >= len(nums) {
			break
		}
		for i := len(matrix[len(matrix)-1]) - 1; i >= 0; i-- {
			nums[index] = matrix[len(matrix)-1][i]
			index++
		}
		if index >= len(nums) {
			break
		}
		matrix = matrix[:len(matrix)-1][:]
		for i := len(matrix) - 1; i >= 0; i-- {
			nums[index] = matrix[i][0]
			matrix[i] = matrix[i][1:]
			index++
		}
	}
	return nums
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiralOrder(matrix))
}
