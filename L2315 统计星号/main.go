package main

import (
	"fmt"
	"strings"
)

func countAsterisks(s string) (ans int) {
	nums := strings.Split(s, "|")
	for i := 0; i < len(nums); i += 2 {
		str := nums[i]
		for j := 0; j < len(str); j++ {
			if str[j] == '*' {
				ans++
			}
		}
	}
	return
}

func main() {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
}
