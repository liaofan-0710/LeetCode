package main

import (
	"fmt"
	"math/bits"
)

func findMaximumNumber(k int64, x int) int64 {
	l, r := int64(1), (k+1)<<x
	for l < r {
		m := (l + r + 1) / 2
		if accumulatedPrice(x, m) > k {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func accumulatedBitPrice(x int, num int64) int64 {
	period := int64(1) << x
	res := period / 2 * (num / period)
	if num%period >= period/2 {
		res += num%period - (period/2 - 1)
	}
	return res
}

func accumulatedPrice(x int, num int64) int64 {
	res := int64(0)
	length := 64 - bits.LeadingZeros64(uint64(num))
	for i := x; i <= length; i += x {
		res += accumulatedBitPrice(i, num)
	}
	return res
}

func main() {
	fmt.Println(findMaximumNumber(9, 1))
}
