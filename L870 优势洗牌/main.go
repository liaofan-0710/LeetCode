package main

import (
	"fmt"
	"sort"
)

func advantageCount1(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	result := make([]int, n)
	idx1 := make([]int, n)
	idx2 := make([]int, n)
	for i := 1; i < n; i++ {
		idx1[i] = i
		idx2[i] = i
	}
	sort.Slice(idx1, func(i, j int) bool { return nums1[idx1[i]] < nums1[idx1[j]] })
	sort.Slice(idx2, func(i, j int) bool { return nums2[idx2[i]] < nums2[idx2[j]] })

	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			result[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			result[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return result
}

func advantageCount(nums1, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool { return nums2[ids[i]] < nums2[ids[j]] })
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[ids[left]] {
			ans[ids[left]] = x // 用下等马比下等马
			left++
		} else {
			ans[ids[right]] = x // 用下等马比上等马
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
