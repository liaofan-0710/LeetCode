package main

import (
	"fmt"
	"math"
	"sort"
)

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	result := 0
	result = int(math.Max(float64(special[0]-bottom), float64(result)))
	result = int(math.Max(float64(top-special[len(special)-1]), float64(result)))
	for i := 1; i < len(special); i++ {
		result = int(math.Max(float64(special[i]-1-special[i-1]), float64(result)))
	}
	return result
}

func main() {
	special := []int{4, 6}
	fmt.Println(maxConsecutive(2, 9, special))
}
