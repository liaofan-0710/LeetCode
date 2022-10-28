package main

import (
	"fmt"
	"math/bits"
)

func oddCells1(m int, n int, indices [][]int) int {
	nums := make([][]int, m)
	ri, ci, res := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, n)
	}
	for i := 0; i < len(indices); i++ {
		ri = indices[i][0]
		ci = indices[i][1]
		for j := 0; j < len(nums[ri]); j++ {
			nums[ri][j]++
		}
		for j := 0; j < len(nums); j++ {
			nums[j][ci]++
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j]%2 != 0 {
				res++
			}
		}
	}
	return res
}

func oddCells(m int, n int, indices [][]int) (ans int) {
	var h, v uint64
	for _, ids := range indices {
		h ^= 1 << ids[0]
		v ^= 1 << ids[1]
	}
	for i := 0; i < m; i++ {
		var x uint64
		if ((h >> i) & 1) == 1 {
			x = ((1 << n) - 1) ^ v
		} else {
			x = v
		}
		ans += bits.OnesCount64(x)
	}
	return
}

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
}
