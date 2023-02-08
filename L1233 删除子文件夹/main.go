package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	ans := []string{}
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return ans
}

func main() {
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
}
