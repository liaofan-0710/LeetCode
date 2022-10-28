package main

import (
	"fmt"
	"strings"
)

func checkOnesSegment1(s string) bool {
	boolean := false
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			boolean = true
		}
		if boolean {
			if s[i] == '1' {
				return false
			}
		}
	}
	return true
}

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	fmt.Println(checkOnesSegment("1111001"))
}
