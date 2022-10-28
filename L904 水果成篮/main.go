package main

import (
	"fmt"
	"runtime/debug"
)

func totalFruit1(fruits []int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

// https://space.bilibili.com/206214
func totalFruit(fruits []int) (ans int) {
	cnt := make([]int, len(fruits))
	left := 0
	c := 0
	for right, x := range fruits {
		if cnt[x] == 0 {
			c++
		}
		cnt[x]++
		for c > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				c--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
func init() { debug.SetGCPercent(-1) }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
