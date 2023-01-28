package main

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
	nums := make([][]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, n)
	}
	for i := 0; i < len(queries); i++ {
		index1 := []int{queries[i][0], queries[i][1]}
		index2 := []int{queries[i][2], queries[i][3]}
		for x := index1[0]; x <= index2[0]; x++ {
			for y := index1[1]; y <= index2[1]; y++ {
				nums[x][y]++
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
}
