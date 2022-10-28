package main

import "fmt"

func moveZeroes3(nums []int) {
	for i, j := 0, 1; j < len(nums); {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				j++
			}
		} else {
			i++
			j = i + 1
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

// 暴力检测
func moveZeroes2(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 优雅解法，教科书级别
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}
