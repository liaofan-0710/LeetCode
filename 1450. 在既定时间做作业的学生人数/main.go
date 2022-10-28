package main

import (
	"fmt"
	"sort"
)

func busyStudent1(startTime []int, endTime []int, queryTime int) int {
	result := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			result++
		}
	}
	return result
}

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	sort.Ints(startTime)
	sort.Ints(endTime)
	return sort.SearchInts(startTime, queryTime+1) - sort.SearchInts(endTime, queryTime)
}

func main() {
	startTime := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	endTime := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	queryTime := 4
	fmt.Println(busyStudent(startTime, endTime, queryTime))
}
