package main

import (
	"fmt"
)

// 第一种：先排序，遍历找出当前数的前一个和后一个都不想等都数字 o(n*log n)
// 第二种：使用map遍历一遍找值等于1都数字 o(n)
// 第三种：进行全与运算 o(n)
// 第四种：开辟一个数组，向后遍历大于二删除，最终返回下标为0都值 (空间o(nums[x]))

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		res |= (sum % 3) << i
	}
	return res
}

// 第二种
func singleNumber2(nums []int) int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; !ok {
			count[nums[i]] = 1
		} else {
			count[nums[i]]++
		}
	}
	for k, v := range count {
		if v == 1 {
			return k
		}
	}
	return nums[0]
}

// 第三种
func singleNumber1(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = ones ^ nums[i] & ^twos
		twos = twos ^ nums[i] & ^ones
	}
	return ones
}

func main() {
	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
}
