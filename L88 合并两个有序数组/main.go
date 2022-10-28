package main

import (
	"fmt"
	"sort"
)

func merge1(nums1 []int, m int, nums2 []int, n int)  {
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	result := make([]int, m+n)
	for i := 0; i < m; i++ {
		result[i] = nums1[i]
	}
	for i, j := m, 0; i < m+n; i++ {
		result[i] = nums2[j]
		j ++
	}
	sort.Ints(result)
	copy(nums1, result)
}

// 错误，逻辑过于多了
func merge2(nums1 []int, m int, nums2 []int, n int)  {
	result := make([]int, m+n)
	k := 0
	for i ,j := 0, 0; k < m+n;{
		if nums1[i] < nums2[j] && nums1[i] != 0 {
			result[k] = nums1[i]
			i ++
			k ++
		}else if nums2[j] <= nums1[i] && nums2[j] != 0 {
			result[k] = nums2[j]
			j ++
			k ++
		}else {
			if nums2[j] == 0 {
				j ++
			}else {
				i ++
			}
		}
	}
	copy(nums1,result)
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	m , n := 3, 3
	merge(nums1,m,nums2,n)
	fmt.Println(nums1)
}