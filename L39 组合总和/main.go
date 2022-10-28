package main

import (
	"fmt"
	_ "net/http/pprof"
	"sort"
)

func combinationSum1(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Sort(sort.IntSlice(candidates))
	var res [][]int
	var path []int
	var backtracking func(res *[][]int, path []int, candidates []int, target int, sum int, idx int)
	backtracking = func(res *[][]int, path []int, candidates []int, target int, sum int, idx int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			*res = append(*res, tmp)
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			backtracking(res, path, candidates, target, sum+candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	backtracking(&res, path, candidates, target, 0, 0)
	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	nums := combinationSum(candidates, target)
	fmt.Println(nums)
}
