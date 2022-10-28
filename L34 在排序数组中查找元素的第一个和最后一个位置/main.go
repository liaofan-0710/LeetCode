package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 8
	nums = searchRange(nums, target)
	fmt.Println(nums[0], nums[1])
}

// 双指针 两头指针一次只挪动一边
func searchRange1(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	j := len(nums)-1
	i1 := true
	j1 := true
	for i := 0; i1 != false || j1 != false; {

		if nums[i] == target && i1 {
			result[0] = i
			i1 = false
		}
		if nums[j] == target && j1 {
			result[1] = j
			j1 = false
		}
		if i1 {
			i ++
		}else if i1 {
			j --
		}

		if i == j{
			break
		}

	}
	return result
}

// 官方题解
func searchRange2(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target + 1) - 1
	return []int{leftmost, rightmost}
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := n-1
	for left <= right{ //二分法查找
		m := (left + right)/2
		if nums[m] == target{
			i := m
			j := m
			for (i>=0 && nums[i] == target) || (j<=n-1 &&nums[j] == target){ //左右指针扩散
				if i>=0 && nums[i] == target{
					i--
				}
				if j<=n-1 && nums[j] == target{
					j++
				}
			}
			return []int{i+1 ,j-1}
		}
		if nums[m] > target{
			right = m-1
		}else{
			left = m+1
		}
	}
	return []int{-1,-1}
}