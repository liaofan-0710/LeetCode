package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	str := strings.Split(s, " ")
	count := math.MinInt
	for i := 0; i < len(str); i++ {
		digit, err := strconv.Atoi(str[i])
		if err == nil {
			if count < digit {
				count = digit
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
}
