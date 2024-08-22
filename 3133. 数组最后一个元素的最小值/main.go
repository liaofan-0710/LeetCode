package main

import (
	"fmt"
	"math/bits"
)

func minEnd(n int, x int) int64 {
	bitCount := 128 - bits.LeadingZeros(uint(n)) - bits.LeadingZeros(uint(x))
	res := int64(x)
	m := int64(n) - 1
	j := 0
	for i := 0; i < bitCount; i++ {
		if res&(1<<i) == 0 {
			if m&(1<<j) != 0 {
				res |= 1 << i
			}
			j++
		}
	}
	return res
}

func main() {
	fmt.Println(minEnd(2, 7))
}
