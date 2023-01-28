package main

import "fmt"

func evaluate(s string, knowledge [][]string) (result string) {
	know := make(map[string]string, len(knowledge))
	for i := 0; i < len(knowledge); i++ {
		know[knowledge[i][0]] = knowledge[i][1]
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			i++
			str := ""
			for s[i] != ')' {
				str += string(s[i])
				i++
			}
			if _, boolean := know[str]; boolean {
				result += know[str]
			} else {
				result += "?"
			}
		} else {
			result += string(s[i])
		}
	}
	return
}

func main() {
	fmt.Println(evaluate("(a)(a)(a)aaa", [][]string{{"a", "yes"}}))
}
