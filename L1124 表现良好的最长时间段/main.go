package main

import "fmt"

func longestWPI(hours []int) int {
	n := len(hours)
	count := make(map[int]int)
	s, res := 0, 0
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			s++
		} else {
			s--
		}
		if s > 0 {
			res = max(res, i+1)
		} else {
			if _, ok := count[s-1]; ok {
				res = max(res, i-count[s-1])
			}
		}
		if _, ok := count[s]; !ok {
			count[s] = i
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
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
}
