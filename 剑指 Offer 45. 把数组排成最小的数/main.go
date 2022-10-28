package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber1(nums []int) string {
	result := ""
	sort.Slice(nums, func(i, j int) bool {
		a, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]))
		b, _ := strconv.Atoi(strconv.Itoa(nums[j]) + strconv.Itoa(nums[i]))
		return a < b
	})
	for i := 0; i < len(nums); i++ {
		result += strconv.Itoa(nums[i])
	}
	return result
}

func minNumber(nums []int) string {
	var strs []string
	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	sort.Slice(strs, func(i, j int) bool { return strs[i]+strs[j] < strs[j]+strs[i] })

	var res string
	for i := 0; i < len(strs); i++ {
		res = res + strs[i]
	}
	return res
}

func main() {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}
