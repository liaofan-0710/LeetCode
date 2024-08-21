package main

import "fmt"

/*
解法1: 常规遍历
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums1); i++ {
		res[0] += isTrue(nums1[i], nums2)
	}
	for i := 0; i < len(nums2); i++ {
		res[1] += isTrue(nums2[i], nums1)
	}
	return res
}
func isTrue(num int, nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return 1
		}
	}
	return 0
}
*/

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	map1, map2 := map[int]bool{}, map[int]bool{}

	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = true
	}

	res := make([]int, 2)
	for i := 0; i < len(nums1); i++ {
		if map2[nums1[i]] {
			res[0]++
		}
	}

	for i := 0; i < len(nums2); i++ {
		if map1[nums2[i]] {
			res[1]++
		}
	}
	return res
}

func main() {
	fmt.Println(findIntersectionValues([]int{2, 3, 2}, []int{1, 2}))
}
