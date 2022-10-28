package main

import "fmt"

func uniquePaths(m int, n int) int {
	return 0
}

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Println(key, &val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
	fmt.Println(uniquePaths(3, 2))
}
