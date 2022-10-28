package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Let's take LeetCode contest"
	result := reverseWords1(str)
	fmt.Println(result)
}

// 暴力更改
func reverseWords(s string) string {
	arrays := strings.Fields(s)
	result := ""
	for i := 0; i < len(arrays)-1; i++ {
		str := revers(arrays[i])
		result += str + " "
	}
	result += revers(arrays[len(arrays)-1])
	return result
}

func revers(s string) string {
	arrays := strings.Split(s,"")
	j := len(arrays)-1
	for i := 0; i < len(arrays)/2; i++ {
		arrays[i], arrays[j] = arrays[j], arrays[i]
		j --
	}
	result := ""
	for i := 0; i < len(arrays); i++ {
		result += arrays[i]
	}
	return result
}

// LeetCode 解法
/*
1、指针f 找到空格
2、双指针翻转单词
3、byte不支持中文，改用rune
*/
func reverseWords1(s string) string {
	str := []rune(s)
	f := 0
	for f < len(str) {
		l := f
		for f < len(str) && str[f] != ' ' {
			f++
		}
		r := f - 1

		for l < r {
			str[l], str[r] = str[r], str[l]
			l++; r--
		}
		f++
	}
	return string(str)
}

/*
总结：每一次找到一个单词的开始下标和结束下标进行局部反转
*/