package main

import "fmt"

func decodeMessage(key string, message string) (result string) {
	count := make(map[byte]string)
	for i, j := 'a', 0; i <= 'z'; {
		if _, ok := count[key[j]]; !ok && key[j] != ' ' {
			count[key[j]] = string(i)
			i++
		}
		j++
	}
	for i := 0; i < len(message); i++ {
		if message[i] != ' ' {
			result += count[message[i]]
		} else {
			result += " "
		}
	}
	return
}

var result = 0

func di(n int) int {
	if n > 100 {
		return 0
	}
	result += n
	return n + di(n+1)
}

func main() {
	res2 := di(1)
	fmt.Println(result) //  出战前的值
	fmt.Println(res2)   // 出栈后的值
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	fmt.Println(decodeMessage(key, message))
}
