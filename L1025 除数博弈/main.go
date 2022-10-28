package main

import "fmt"

func divisorGame(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(divisorGame(5))

}
