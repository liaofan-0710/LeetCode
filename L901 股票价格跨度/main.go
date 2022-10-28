package main

import (
	"fmt"
	"math"
)

type StockSpanner struct {
	stack [][2]int
	idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, -1}
}

func (this *StockSpanner) Next(price int) int {
	this.idx++
	for price >= this.stack[len(this.stack)-1][1] {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{this.idx, price})
	return this.idx - this.stack[len(this.stack)-2][0]
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
