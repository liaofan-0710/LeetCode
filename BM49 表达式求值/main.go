package main

import (
	"fmt"
)

func solve(s string) int {
	i := 0
	var helper func(string) int
	helper = func(s string) int {
		sign := byte('+')
		num := 0
		bytes := []byte(s)
		stack := []int{}
		for i < len(bytes) {
			b := s[i]
			i++
			if isDigit(b) {
				num = num*10 + int(b-'0')
			}
			if b == '(' {
				num = helper(s)
			}
			if !isDigit(b) && b != ' ' || i == len(bytes) {
				switch sign {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				case '*':
					stack[len(stack)-1] *= num
				case '/':
					stack[len(stack)-1] /= num
				}
				sign = b
				num = 0
			}
			if b == ')' {
				break
			}
		}
		return sum(stack)
	}
	return helper(s)
}
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
func sum(stack []int) int {
	s := 0
	for _, num := range stack {
		s += num
	}
	return s
}

func main() {
	fmt.Println(solve("(2*(3-4))*5"))
}
