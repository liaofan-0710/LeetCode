package main

import (
	"fmt"
	"sort"
)

// map计数
func majorityElement1(nums []int) int {
	count := make(map[int]int, 0)
	res := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; !ok {
			count[nums[i]] = 1
		} else {
			count[nums[i]]++
		}
	}
	for k, v := range count {
		if v > len(nums)/2 {
			res = k
			break
		}
	}
	return res
}

// 排序找取记录数字，判断是否大于长度/2
func majorityElement2(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	num := []int{nums[0], 1}
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			count = 1
		}
		if num[1] <= count {
			num[0] = nums[i-1]
			num[1] = count
		}
	}
	return num[0]
}

// 随机选择数据(超时)
func majorityElement(nums []int) int {
	count, target, i := 0, len(nums)/2, 0
	zan, j := nums[0], 0
	for target >= count {
		i++
		if i >= len(nums) {
			zan = nums[j]
			i = 0
			count = 0
			j++
		}
		if nums[i] == zan {
			count++
		}
	}
	return nums[i]
}

func main() {
	nums := []int{1}
	fmt.Println(majorityElement(nums))
}
