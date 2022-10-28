package main

import (
	"fmt"
	"strings"
)

func areAlmostEqual1(s1 string, s2 string) bool {
	if strings.Compare(s1, s2) == 0 {
		return true
	}
	ans := make(map[string]string)
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if len(ans) > 0 && ans[string(s2[i])] != string(s1[i]) {
				return false
			}
			if _, ok := ans[string(s1[i])]; !ok {
				ans[string(s1[i])] = string(s2[i])
			}
		}
		if count > 2 {
			return false
		}
	}
	if len(ans) == 1 {
		return false
	}
	return true
}

func areAlmostEqual(s1, s2 string) bool {
	i, j := -1, -1
	for k := 0; k < len(s1); k++ {
		if s1[k] != s2[k] {
			if i < 0 {
				i = k
			} else if j < 0 {
				j = k
			} else {
				return false
			}

		}
	}
	return i < 0 || j >= 0 && s1[i] == s2[j] && s1[j] == s2[i]
}

func main() {
	fmt.Println(areAlmostEqual("caa", "aaz"))
}
