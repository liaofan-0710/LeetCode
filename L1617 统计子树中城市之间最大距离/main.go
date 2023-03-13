package main

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	// 建树
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1 // 编号改为从 0 开始
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 求树的直径
	var inSet, vis [15]bool
	var diameter int
	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] && inSet[y] {
				ml := dfs(y) + 1
				diameter = max(diameter, maxLen+ml)
				maxLen = max(maxLen, ml)
			}
		}
		return
	}

	ans := make([]int, n-1)
	var f func(int)
	f = func(i int) {
		if i == n {
			for v, b := range inSet {
				if b {
					vis, diameter = [15]bool{}, 0
					dfs(v)
					break
				}
			}
			if diameter > 0 && vis == inSet {
				ans[diameter-1]++
			}
			return
		}

		// 不选城市 i
		f(i + 1)

		// 选城市 i
		inSet[i] = true
		f(i + 1)
		inSet[i] = false // 恢复现场
	}
	f(0)
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

}
