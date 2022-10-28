package main

import "fmt"

var testList = make([]int, 10)

// 获取中间值
func partition(nums []int, first, last int) int {
	pivotValue := nums[first]
	leftMark := first + 1
	rightMark := last
	done := false
	for !done {
		for leftMark <= rightMark && nums[leftMark] <= pivotValue {
			leftMark++
		}
		for rightMark >= leftMark && nums[rightMark] >= pivotValue {
			rightMark--
		}
		if rightMark < leftMark {
			done = true
		} else {
			temp := nums[leftMark]
			nums[leftMark] = nums[rightMark]
			nums[rightMark] = temp
		}
	}
	// 中值就位
	temp := nums[first]
	nums[first] = nums[rightMark]
	nums[rightMark] = temp
	return rightMark
}

func quickHelper(nums []int, first, last int) {
	if first < last {
		// 获取分裂点
		splitPoint := partition(nums, first, last)
		// 通过分裂点，将左右两个子列表分别快排
		quickHelper(nums, first, splitPoint-1)
		quickHelper(nums, splitPoint+1, last)
	}
}

func quickSort(nums []int) {
	// 分裂准备
	quickHelper(nums, 0, len(nums)-1)
}

func main() {
	testList = []int{9, 7, 2, 323, 43, 5, 6, 66, 45}
	quickSort(testList)
	fmt.Println(testList)
}
