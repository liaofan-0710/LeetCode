package main

import "fmt"

func findContinuousSequence1(target int) [][]int {
	i, j := 1, 1
	sum := 0
	nums := make([][]int, 0)
	num := make([]int, target-1)
	for k := 0; k < target-1; k++ {
		num[k] = k + 1
	}
	for i <= target/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			nums = append(nums, num[i-1:j-1])
			sum -= i
			i++
		}
	}
	return nums
}

func findContinuousSequence(target int) [][]int {
	var res [][]int
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2
		switch {
		case sum == target: // Save res parts
			var numBuffer []int
			for i := l; i <= r; i++ {
				numBuffer = append(numBuffer, i)
			}
			res = append(res, numBuffer)
			l += 1
			r += 1
		case sum < target:
			r += 1
		case sum > target:
			l += 1
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(15))
}
