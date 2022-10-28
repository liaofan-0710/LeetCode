package main

import "fmt"

// 动态规划
// 状态：左括号：右括号
func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}

		}
	}
	return dp[n]
}

func main() {
	fmt.Println(generateParenthesis(3))
}
