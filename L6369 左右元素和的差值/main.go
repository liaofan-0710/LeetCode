package main

import "fmt"

func leftRigthDifference(nums []int) []int {
	n := len(nums)
	leftSum, rightSum, answer := make([]int, n), make([]int, n), make([]int, n)
	for i, j := 0, n-1; i < n-1; {
		leftSum[i+1] = leftSum[i] + nums[i]
		rightSum[j-1] = rightSum[j] + nums[j]
		i++
		j--
	}
	for i := 0; i < n; i++ {
		if leftSum[i]-rightSum[i] > 0 {
			answer[i] = leftSum[i] - rightSum[i]
		} else {
			answer[i] = (leftSum[i] - rightSum[i]) * -1
		}
	}
	return answer
}

func main() {
	fmt.Println(leftRigthDifference([]int{1}))
}
