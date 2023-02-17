package main

import (
	"fmt"
	"sort"
)

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	sort.Ints(nums)
	min1 := make([]int, 0)
	max1 := make([]int, 0)
	if sum%2 != 0 {
		return false
	}
	middle := sum / 2
	for i := 0; i < len(nums); i++ {
		if nums[i] < middle {
			min1 = append(min1, nums[i])
		} else {
			max1 = append(max1, nums[i])
		}
	}
	res := combinationSum2(min1, middle)
	if len(res) > 0 {
		return true
	}
	return false
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
