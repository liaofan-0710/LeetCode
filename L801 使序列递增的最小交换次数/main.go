package main

import (
	"fmt"
)

func minSwap(nums1, nums2 []int) int {
	n, notSwap, swap := len(nums1), 0, 1
	for i := 1; i < n; i++ {
		ns, s := n, n // 答案不会超过 n，故初始化成 n 方便后面取 min
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			ns, s = notSwap, swap+1
		}
		if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
			ns, s = min(ns, swap), min(s, notSwap+1)
		}
		notSwap, swap = ns, s
	}
	return min(notSwap, swap)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}))
}
