package main

func main() {
	var nums = []int{1,3,5,6}
	target := 7
	result := search(nums, target)
	print(result)
}

// 普通写法  遍历一遍暴力查找
func search(nums []int, target int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
		}
	}
	return result
}

// 二分查找
/*
思想：与中间值进行对比
1. 如果target比中间值小则向左折半查找，
2. 如果target比中间值大则向右折半查找
3. 每一次查找直到找到返回否则返回-1
*/
func search1(nums []int, target int) int {
	length := len(nums)
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
