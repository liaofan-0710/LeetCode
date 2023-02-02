package main

import "fmt"

func numTrees1(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func numTrees(n int) int {

	if n == 0 || n == 1 {
		return 1
	}
	count := 0
	for i := 1; i <= n; i++ {
		left := numTrees(i - 1)
		right := numTrees(n - i)
		count += left * right
	}
	return count
}

func main() {
	fmt.Println(numTrees(5))
}
