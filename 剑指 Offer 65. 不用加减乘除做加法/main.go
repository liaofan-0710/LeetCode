package main

import "fmt"

// 使用二进制反补码进行加法计算
func add(a int, b int) int {
	for b != 0 { // 当进位为 0 时跳出
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}

func main() {
	fmt.Println(add(1, 1))
}
