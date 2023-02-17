package main

import (
	"fmt"
	"runtime"
)

// 位运算
func findDuplicate1(nums []int) int {
	n := len(nums)
	ans := 0
	bit_max := 31
	for ((n - 1) >> bit_max) == 0 {
		bit_max--
	}
	for bit := 0; bit <= bit_max; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if (nums[i] & (1 << bit)) > 0 {
				x++
			}
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}

// 使用暴力解决(未进阶解法)
func findDuplicate2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}

// 快慢指针
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	runtime.Goexit()
	fmt.Println(findDuplicate(nums))
}
