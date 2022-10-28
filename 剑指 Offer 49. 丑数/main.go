package main

import "fmt"

// 题解一：前缀和
// 题解二：循环遍历

// 循环遍历(超时)
func nthUglyNumber1(n int) int {
	if n == 0 {
		return n
	}
	j := 1
	for i := 0; i < n; {
		k := j
		for k > 1 {
			if k%2 == 0 {
				k /= 2
			} else if k%5 == 0 {
				k /= 5
			} else if k%3 == 0 {
				k /= 3
			} else {
				break
			}
		}
		if k == 1 {
			i++
		}
		j++
	}
	return j - 1
}

// 动态规划
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
