package main

import "fmt"

// 第一种：遍历 1. 找到最小的值 2. 以最小值为中间点找右边最大的值  3. 进行相减获得最大值
// 第二种：一次遍历，不断更新最小值，进行相减
// 第二种：dp 状态： 1.买/不买  2. 卖/不卖
func maxProfit1(prices []int) int {
	// 每一次假设下标0就是最小的然后找到右边比它大的进行相减，不断更新max
	min, result := 0, 0
	for len(prices) > 0 {
		min = prices[0]
		for i := 1; i < len(prices); i++ {
			if prices[i] > min {
				result = maxS(result, prices[i]-min)
			}
		}
		prices = prices[1:]
	}
	return result
}

// 一次遍历
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, result := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			result = maxS(result, prices[i]-min)
		} else if min > prices[i] {
			min = prices[i]
		}
	}
	return result
}

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, result := prices[0], 0
	for i := 1; i < len(prices); i++ {
		min = minS(min, prices[i])
		result = maxS(result, prices[i]-min)
	}
	return result
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := prices[0]
	for _, v := range prices {
		if v > min {
			max = maxS(max, v-min)
		} else {
			min = v
		}
	}
	return max
}

func maxS(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minS(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
