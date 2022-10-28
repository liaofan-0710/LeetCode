package main

import (
	"fmt"
	"strings"
)

func canChange(start, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	j := 0
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' {
				j++
			}
			if i != j && c == 'L' != (i > j) {
				return false
			}
			j++
		}
	}
	return true
}

func main() {
	fmt.Println(canChange("_L__R__R_", "L______RR"))
}
