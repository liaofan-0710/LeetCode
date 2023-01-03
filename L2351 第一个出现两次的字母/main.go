package main

import "fmt"

// 第一种使用Map
func repeatedCharacter1(s string) byte {
	count := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if count[s[i]] {
			return s[i]
		}
		count[s[i]] = true
	}
	return 0
}

// 第二种 数组存储
// 解释：因为字母只有二十六个，所以采用下表加储存值的方法进行记录
func repeatedCharacter2(s string) byte {
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] >= 1 {
			return s[i]
		}
		count[s[i]-'a']++
	}
	return 0
}

// 第三种 力扣题解状态压缩
func repeatedCharacter(s string) byte {
	seen := 0
	for _, c := range s {
		if seen>>(c-'a')&1 > 0 {
			return byte(c)
		}
		seen |= 1 << (c - 'a')
	}
	return 0 // impossible
}

func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}
