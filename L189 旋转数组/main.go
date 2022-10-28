package main

import "fmt"

func main() {
	nums := []int{-1,2,30}
	k := 10
	rotate(nums, k)
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
}
func rotate(nums []int, k int)  {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
func rotate1(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
