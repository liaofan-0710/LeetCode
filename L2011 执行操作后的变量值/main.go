package main

import "fmt"

func finalValueAfterOperations(operations []string) (ans int) {
	for _, v := range operations {
		if v == "++X" || v == "X++" {
			ans += 1
		} else {
			ans -= 1
		}
	}
	return
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))
}
