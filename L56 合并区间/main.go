package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, []int{intervals[0][0], intervals[0][1]})
	j := 1
	for i := 1; i < len(intervals); i++ {
		ans := []int{intervals[i][0], intervals[i][1]}
		if res[j-1][1] >= intervals[i][0] {
			res[j-1][1] = max(res[j-1][1], intervals[i][1])
		} else {
			res = append(res, ans)
			j++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {0, 6}}))
}
