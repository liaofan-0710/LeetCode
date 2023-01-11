package main

import (
	"fmt"
)

func digitCount(num string) bool {
	nums := [10]int{}
	for i := 0; i < len(num); i++ {
		nums[num[i]-48]++
	}
	for i := 0; i < len(num); i++ {
		if int(num[i])-48 != nums[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(digitCount("030"))
}
