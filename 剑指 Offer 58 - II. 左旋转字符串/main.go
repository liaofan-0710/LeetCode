package main

import (
	"fmt"
)

func reverseLeftWords(s string, n int) string {
	return s[n:len(s)] + s[:n]
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}
