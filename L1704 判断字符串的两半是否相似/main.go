package main

import (
	"fmt"
	"strings"
)

func halvesAreAlike1(s string) bool {
	count := map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
		"A": 1,
		"E": 1,
		"I": 1,
		"O": 1,
		"U": 1,
	}
	a, b := 0, 0
	for i := 0; i < len(s)/2; i++ {
		if count[string(s[i])] == 1 {
			a++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if count[string(s[i])] == 1 {
			b++
		}
	}
	if a != b {
		return false
	}
	return true
}

func halvesAreAlike(s string) bool {
	cnt := 0
	for i := 0; i < len(s)/2; i++ {
		if strings.Contains("aeiouAEIOU", string(s[i])) {
			cnt++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if strings.Contains("aeiouAEIOU", string(s[i])) {
			cnt--
		}
	}
	return cnt == 0
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
