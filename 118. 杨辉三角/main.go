package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	return result
}

func main() {
	fmt.Println(generate(5))
}
