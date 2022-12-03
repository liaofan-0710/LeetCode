package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func secondHighest1(s string) int {
	res := make([]int, 0)
	ans := map[int]int{}
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' {
			count, _ := strconv.Atoi(string(s[i]))
			if _, ok := ans[count]; !ok {
				res = append(res, count)
				ans[count]++
			}
		}
	}
	sort.Ints(res)
	if len(res) > 1 && res[len(res)-1] > res[len(res)-2] {
		return res[len(res)-2]
	}
	return -1
}

func secondHighest(s string) int {
	first, second := -1, -1
	for _, c := range s {
		if unicode.IsDigit(c) {
			num := int(c - '0')
			if num > first {
				second = first
				first = num
			} else if second < num && num < first {
				second = num
			}
		}
	}
	return second
}

func main() {
	fmt.Println(secondHighest("ck077"))
}
