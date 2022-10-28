package main

import "fmt"

func main()  {
	nums := []int{1,2,3,1}
	result := containsDuplicate(nums)
	fmt.Println(result)
}


func containsDuplicate(nums []int) bool {
	mapNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mapNums[nums[i]] ++
		if mapNums[nums[i]] > 1 {
			return true
		}
	}
	return false
}