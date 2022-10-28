package main

import (
	"fmt"
	"sort"
)

// 第一种：先判断少几张牌没有连续或者已经连续，最终计算0有多少个进行比对返回结果
func isStraight1(nums []int) bool {
	sort.Sort(sort.IntSlice(nums))
	count1 := 0
	count2 := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == 0 {
			count1++
			continue
		}
		if nums[i-1]+1 != nums[i] {
			if nums[i-1] == nums[i] && nums[i] != 0 {
				return false
			}
			if nums[i] != 0 {
				count2 += nums[i] - nums[i-1] - 1
			}
		}
	}
	return count1 >= count2
}

// 第二种：
func isStraight(nums []int) bool {
	sort.Sort(sort.IntSlice(nums))
	count := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			count++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[count] < 5 // 最大牌 - 最小牌 < 5 则可构成顺子
}

func main() {
	fmt.Println(isStraight([]int{11, 0, 9, 0, 0}))
}
