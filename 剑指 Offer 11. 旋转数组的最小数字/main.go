package main

import "fmt"

func minArray(numbers []int) int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		pivot := i + (j-i)/2
		if numbers[pivot] < numbers[j] {
			j = pivot
		} else if numbers[pivot] > numbers[j] {
			i = pivot + 1
		} else {
			j--
		}
	}
	return numbers[j]
}

func main() {
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
