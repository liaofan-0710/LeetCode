package main

import "fmt"

func prefixCount(words []string, pref string) (result int) {
	n := len(pref)
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= n && pref == words[i][:n] {
			result++
		}
	}
	return
}

func main() {
	words := []string{"sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "vbx", "fsi", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "gqira", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh", "sxyjellhlh"}
	pref := "sxyjellhlh"
	fmt.Println(prefixCount(words, pref))
}
