package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	res := make([]int, int(math.Pow10(n)-1))
	for i := 0; i < len(res); i++ {
		res[i] = i + 1
	}
	return res
}

// 考虑大数情况

func main() {
	fmt.Println(printNumbers(1))
}
