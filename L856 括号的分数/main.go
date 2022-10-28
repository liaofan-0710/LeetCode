package main

import "fmt"

func scoreOfParentheses(s string) (ans int) {
	bal := 0
	for i, c := range s {
		if c == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				ans += 1 << bal
			}
		}
	}
	return
}

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
}
