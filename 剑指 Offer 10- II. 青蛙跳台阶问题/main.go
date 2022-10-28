package main

import "fmt"

func numWays1(n int) int {
	pre, cur, res := 1, 1, 0
	for i := 0; i < n; i++ {
		res = (pre + cur) % 1000000007
		pre = cur
		cur = res
	}
	return pre
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	pre := 1
	cur := 2
	for i := 3; i <= n; i++ {
		res := (pre + cur) % 1000000007
		pre = cur
		cur = res
	}
	return cur
}

func main() {
	fmt.Println(numWays(7))
}
