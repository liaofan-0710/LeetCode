package main

import "fmt"

// 第一种：假设第一个就是第一次出现的数字，否则继续向后查找(这个方法不可行)
// 复杂度分析：
// 第二种：使用map进行计数，只要出现一次的字母，之后遍历比对，只要出现一次的相等就返回
// 复杂度分析:

// 思路：使用map记录，之后重新循环的时候进行判断值是否为1
func firstUniqChar1(s string) byte {
	count := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := count[s[i]]; !ok {
			count[s[i]] = 1
		} else {
			count[s[i]]++
		}
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}

func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
}
