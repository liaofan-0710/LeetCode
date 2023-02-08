package main

import "fmt"

//type RecentCounter struct {
//	nums []int
//}
//
//func Constructor() RecentCounter {
//	return RecentCounter{}
//}
//
//func (this *RecentCounter) Ping(t int) int {
//	res := 0
//	this.nums = append(this.nums, t)
//	for _, v := range this.nums {
//		if t-3000 <= v && t >= v {
//			res++
//		}
//	}
//	return res
//}

type RecentCounter []int

func Constructor() (_ RecentCounter) { return }

func (q *RecentCounter) Ping(t int) int {
	*q = append(*q, t)
	for (*q)[0] < t-3000 {
		*q = (*q)[1:]
	}
	return len(*q)
}

func main() {
	recentCounter := Constructor()        // nil
	fmt.Println(recentCounter.Ping(1))    // 1
	fmt.Println(recentCounter.Ping(100))  // 2
	fmt.Println(recentCounter.Ping(3001)) // 3
	fmt.Println(recentCounter.Ping(3002)) // 3
}
