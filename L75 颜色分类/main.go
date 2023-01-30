package main

import "fmt"

func sortColors1(nums []int) {
	ans := 0
	for i, count := 0, 0; i < len(nums); i++ {
		if ans == 0 && nums[i] == 0 {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
		if ans == 1 && nums[i] == 1 {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
		if i == len(nums)-1 {
			i = count - 1
			ans++
		}
		if ans == 2 {
			break
		}
	}
}

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
