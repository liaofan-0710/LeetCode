package main

// 两边放置房子的个数是一样的
// 所以答案是 count(对面) * count(这边)
// 也可以是 count(这边) ^ 2
func countHousePlacements(n int) int {
	// mod 1000000007
	count := count(n)
	return int(count * count % 1000000007)
}

// 由于不能隔壁放置, 所以 f(x) = f(x - 1) + f(x - 2)
// 用动态规划递推
func count(n int) int64 {
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}
	if n == 3 {
		return 5
	}

	dp := make([]int64, n+1)
	dp[1] = 2
	dp[2] = 3
	for i := 3; i < len(dp); i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

func main() {

}
