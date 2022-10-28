package main

import (
	"fmt"
	"strings"
)

func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cns1 := make(map[string]string)
	cns2 := make(map[string]string)
	for i := 0; i < len(s); i++ {
		if _, ok := cns2[string(t[i])]; ok {
			if cns2[string(t[i])] == string(s[i]) {
				continue
			} else {
				return false
			}
		} else {
			cns2[string(t[i])] = string(s[i])
		}
		if _, ok := cns1[string(s[i])]; ok {
			if cns1[string(s[i])] == string(t[i]) {
				continue
			} else {
				return false
			}
		} else {
			cns1[string(s[i])] = string(t[i])
		}
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	s1 := map[byte]byte{}
	s2 := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		x, y := s[i], t[i]
		if s1[x] > 0 && s1[x] != y || s2[y] > 0 && s2[y] != x {
			return false
		}
		s1[x] = y
		s2[y] = x
	}
	return true
}

func isIsomorphic3(s string, t string) bool {
	return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}
func isIsomorphicHelper(s string, t string) bool {
	n := len(s)
	ans := map[byte]byte{}
	for i := 0; i < n; i++ {
		if _, ok := ans[s[i]]; ok {
			if ans[s[i]] != t[i] {
				return false
			}
		} else {
			ans[s[i]] = t[i]
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i, ch := range s {
		if strings.IndexRune(s, ch) != strings.IndexByte(t, t[i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
