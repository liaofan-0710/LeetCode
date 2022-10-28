package main

import "fmt"

// 递归
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	fmt.Print(n, "  ")
	return (fib(n-1) + fib(n-2)) % 1000000007
}

// 记忆化搜索，去重
func fib2(n int) int {
	nums := make([]int, n+1)
	for i := 0; i < len(nums); i++ {
		nums[i] = -1
	}
	return dfs(n, nums)
}

func dfs(n int, nums []int) int {
	if n <= 1 {
		return n
	}
	if nums[n] != -1 { // 如果上次已经计算过了直接返回
		return nums[n]
	}
	nums[n] = (dfs(n-1, nums) + dfs(n-2, nums)) % 1000000007
	fmt.Println(nums)
	return nums[n]
}

func fib3(n int) int {
	nums := make([]int, n+1)
	if n <= 1 {
		return n
	}
	nums[0] = 0
	nums[1] = 1
	for i := 2; i < len(nums); i++ {
		nums[i] = (nums[i-1] + nums[i-2]) % 1000000007
	}
	return nums[n]
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	prev := 0
	curr := 1
	for i := 2; i <= n; i++ {
		res := (prev + curr) % 1000000007
		prev = curr
		curr = res
	}
	return curr
}

func main() {
	fmt.Println(fib(10))
}
