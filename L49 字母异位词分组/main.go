package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	count := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		coust := []byte(strs[i])
		sort.Slice(coust, func(i, j int) bool {
			return coust[i] < coust[j]
		})
		str := string(coust)
		count[str] = append(count[str], strs[i])
	}
	res := make([][]string, len(count))
	i := 0
	for _, strings := range count {
		res[i] = strings
		i++
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
