package main

import "fmt"

func nthMagicalNumber(n, a, b int) int {
	lcm := a / gcd(a, b) * b
	left, right := 0, min(a, b)*n // 开区间 (left, right)
	for left+1 < right {          // 开区间不为空
		mid := left + (right-left)/2
		if mid/a+mid/b-mid/lcm >= n {
			right = mid // 范围缩小到 (left, mid)
		} else {
			left = mid // 范围缩小到 (mid, right)
		}
	}
	return right % (1e9 + 7)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	fmt.Println(nthMagicalNumber(1000000000, 40000, 40000))
}
