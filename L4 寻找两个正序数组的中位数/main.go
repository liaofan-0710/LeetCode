package main

import (
	"fmt"
	"strconv"
)

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	i1, i2, boolean := 0, 0, false
	if n%2 == 0 {
		boolean = true
	}
	nums := make([]int, n)
	for i := 0; i < n/2; i++ {
		if i1 >= len(nums1) {
			nums[i] = nums2[i2]
			i2++
			continue
		} else if i2 >= len(nums2) {
			nums[i] = nums1[i1]
			i1++
			continue
		}
		if nums1[i1] < nums2[i2] {
			nums[i] = nums1[i1]
			i1++
		} else if nums1[i1] > nums2[i2] {
			nums[i] = nums2[i2]
			i2++
		} else {
			nums[i] = nums2[i2]
			i2++
			i++
			nums[i] = nums1[i1]
			i1++
		}
	}
	if !boolean {
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nums[n/2])), 64)
		return value
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nums[n/2]+nums[n/2-1])/2), 64)
	return value
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 4, 4}))
}
