package main

import (
	"fmt"
	"strconv"
)

func divisibilityArray(word string, m int) []int {
	n := len(word)
	div := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		num := 0
		if count != 0 {
			num, _ = strconv.Atoi(strconv.Itoa(count) + string(word[i]))
		} else {
			num, _ = strconv.Atoi(string(word[i]))
		}
		count = num % m
		if count == 0 {
			div[i] = 1
		} else {
			div[i] = 0
		}
	}
	return div
}

func main() {
	fmt.Println(divisibilityArray("91221181269244172125025075166510211202115152121212341281327",
		21))
}
