package main

import "fmt"

func minimumDeletions(s string) int {
	left := 0
	right := 0
	for _, c := range s {
		if c == 'a' {
			right++
		}
	}
	res := right
	for _, c := range s {
		if c == 'a' {
			right--
		} else {
			left++
		}
		res = min(res, left+right)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumDeletions("aababbab"))
}
