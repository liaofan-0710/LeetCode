package main

import "fmt"

func reverseLeftWords1(s string, n int) string {
	return s[n:] + s[:n]
}

func reverseLeftWords(s string, n int) string {
	res := ""
	for i := n; i < len(s); i++ {
		res += string(s[i])
	}
	for i := 0; i < n; i++ {
		res += string(s[i])
	}
	return res
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}
