package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i int, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}

func main() {
	nums := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder(nums))
}
