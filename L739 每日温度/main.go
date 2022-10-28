package main

import "fmt"

func dailyTemperatures1(temperatures []int) []int {
	nums := make([]int, len(temperatures))
	nums[len(temperatures)-1] = 0
	for i := len(temperatures) - 1; i > 0; i-- {
		if temperatures[i] > temperatures[i-1] {
			nums[i-1] = 1
		}
	}
	for i := 0; i < len(temperatures)-1; i++ {
		if nums[i] == 0 {
			count := 0
			for j := i; j < len(temperatures); j++ {
				if temperatures[i] < temperatures[j] {
					break
				}
				count++
				if j == len(temperatures)-1 && temperatures[i] >= temperatures[j] {
					count = 0
				}
			}
			nums[i] = count
		}
	}
	return nums
}

// 暴力
func dailyTemperatures2(t []int) []int {
	var res []int
	for i := 0; i < len(t)-1; i++ {
		j := i + 1
		for ; j < len(t); j++ {
			// 如果之后出现更高，说明找到了
			if t[j] > t[i] {
				res = append(res, j-i)
				break
			}
		}
		// 找完了都没找到
		if j == len(t) {
			res = append(res, 0)
		}
	}
	// 最后一个肯定是 0
	return append(res, 0)
}

// 单调递减栈
func dailyTemperatures(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	temperatures := []int{34, 80, 80, 80, 34, 80, 80, 80, 34, 34}
	fmt.Println(dailyTemperatures(temperatures))
}
