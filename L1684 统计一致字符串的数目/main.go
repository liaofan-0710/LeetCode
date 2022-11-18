package main

import (
	"fmt"
	"strings"
)

func countConsistentStrings1(allowed string, words []string) (ans int) {
	for i := 0; i < len(words); i++ {
		str := words[i]
		boolean := true
		for j := 0; j < len(str); j++ {
			if !strings.Contains(allowed, string(str[j])) {
				boolean = false
			}
		}
		if boolean {
			ans++
		}
	}
	return
}

func countConsistentStrings(allowed string, words []string) int {
	valid := make([]bool, 26)
	for i := range allowed {
		valid[allowed[i]-'a'] = true
	}

	res := 0
	for i := range words {
		s := words[i]
		yes := true
		for k := range s {
			if !valid[s[k]-'a'] {
				yes = false
				break
			}
		}
		if yes {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
}
