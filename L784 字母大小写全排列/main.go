package main

import (
	"fmt"
	"unicode"
)

// 广度优先搜索
func letterCasePermutation(s string) (ans []string) {
	q := []string{""}
	for len(q) > 0 {
		cur := q[0]
		pos := len(cur)
		if pos == len(s) {
			ans = append(ans, cur)
			q = q[1:]
		} else {
			if unicode.IsLetter(rune(s[pos])) {
				q = append(q, cur+string(s[pos]^32))
			}
			q[0] += string(s[pos])
		}
	}
	return
}

// 回溯法
func letterCasePermutation1(s string) (ans []string) {
	var dfs func([]byte, int)
	dfs = func(s []byte, pos int) {
		for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
			pos++
		}
		if pos == len(s) {
			ans = append(ans, string(s))
			return
		}
		dfs(s, pos+1)
		s[pos] ^= 32
		dfs(s, pos+1)
		s[pos] ^= 32
	}
	dfs([]byte(s), 0)
	return
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
