package main

import (
	"fmt"
	"strings"
)

func maxRepeating(sequence string, word string) int {
	for k := len(sequence) / len(word); k > 0; k-- {
		if strings.Contains(sequence, strings.Repeat(word, k)) {
			return k
		}
	}
	return 0
}

func main() {
	str1 := "aaabaaaabaaabaaaabaaaabaaaabaaaaba"
	str2 := "aaaba"
	fmt.Println(maxRepeating(str1, str2))
}
