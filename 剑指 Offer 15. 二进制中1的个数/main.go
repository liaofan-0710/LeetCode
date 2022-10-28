package main

import "fmt"

func hammingWeight1(num uint32) int {
	res := 0
	for num != 0 {
		if num%2 != 0 {
			res++
		}
		num /= 2
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	for ; num != 0; num &= num - 1 {
		res++
	}
	return res
}

func main() {
	fmt.Println(hammingWeight(11))
}
