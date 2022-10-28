package main

import "fmt"

// 第一种：遍历一遍 时间复杂度(n 2次方) 空间复杂度 (O(n)) ！！！超出时间限制
func constructArr1(a []int) []int {
	nums := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		nums[i] = 1
		for j := 0; j < len(a); j++ {
			if i != j {
				nums[i] *= a[j]
			}
		}
	}
	return nums
}

// 第二种 左右乘积列表
func constructArr(a []int) []int {
	length := len(a)
	if length == 0 {
		return a
	}
	// L 和 R 分别表示左右两侧的乘积列表
	L, R, answer := make([]int, length), make([]int, length), make([]int, length)

	// L[i] 为索引 i 左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = a[i-1] * L[i-1]
	}

	// R[i] 为索引 i 右侧所有元素的乘积
	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = a[i+1] * R[i+1]
	}

	// 对于索引 i，除 a[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

func productExceptSelf(a []int) []int {
	length := len(a)
	if length == 0 {
		return a
	}
	answer := make([]int, length)

	// answer[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = a[i-1] * answer[i-1]
	}

	// R 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	R := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		answer[i] = answer[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= a[i]
	}
	return answer
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 0, 4, 5}))
}
