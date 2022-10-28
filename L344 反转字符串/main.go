package main

import "fmt"

func main() {
	s := []byte{'H','a','n','n','a','h'}
	reverseString(s)
	fmt.Print(string(s))
}

func reverseString(s []byte)  {
	j := len(s)-1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j --
	}
}