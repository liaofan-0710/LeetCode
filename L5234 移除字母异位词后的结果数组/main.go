package main

import (
	"fmt"
	"sort"
)

func removeAnagrams(words []string) []string {
	m := map[string]string{}
	var result []string
	for i := 0; i < len(words); i++ {
		str := []byte(words[i])
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		if _, ok := m[string(str)]; ok == false {
			m[string(str)] = words[i]
			result = append(result, words[i])
		}
	}
	return result
}

func main() {
	nums := []string{"a", "b", "a"}
	fmt.Println(removeAnagrams(nums))
}
