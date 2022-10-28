package main

import (
	"fmt"
	"strings"
)

func reverseWords1(s string) string {
	res := ""
	str := strings.Split(s, " ")
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != "" {
			res += str[i] + " "
		}
	}
	if len(res) == 0 {
		return ""
	}
	return res[:len(res)-1]
}

func reverseWords(s string) string {
	var res string
	var i = len(s) - 1
	var j = i
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res += s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}

func main() {
	fmt.Println(reverseWords("  	"))
}
