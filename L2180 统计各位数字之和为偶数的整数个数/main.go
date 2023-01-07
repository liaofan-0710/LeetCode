package main

import (
	"fmt"
	"strconv"
)

// 卡边境，取巧
func countEven1(num int) (res int) {
	for i := 2; i <= num; i++ {
		if len(strconv.Itoa(i)) == 1 && i%2 == 0 {
			res++
		} else if len(strconv.Itoa(i)) == 2 && ((i%10)+(i/10))%2 == 0 {
			res++
		} else if len(strconv.Itoa(i)) == 3 && ((i%10)+(i%100/10)+(i/100))%2 == 0 {
			res++
		}
	}
	return
}

func countEven(num int) (ans int) {
	for i := 1; i <= num; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		if sum%2 == 0 {
			ans++
		}
	}
	return
}

func main() {
	fmt.Println(countEven(30))
}
