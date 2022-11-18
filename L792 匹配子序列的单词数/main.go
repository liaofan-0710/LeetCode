package main

import (
	"fmt"
	"sort"
)

func numMatchingSubseq1(s string, words []string) int {
	pos := [26][]int{}
	for i, c := range s {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	ans := len(words)
	for _, w := range words {
		if len(w) > len(s) {
			ans--
			continue
		}
		p := -1
		for _, c := range w {
			ps := pos[c-'a']
			j := sort.SearchInts(ps, p+1)
			if j == len(ps) {
				ans--
				break
			}
			p = ps[j]
		}
	}
	return ans
}

func numMatchingSubseq(s string, words []string) (ans int) {
	type pair struct{ i, j int }
	ps := [26][]pair{}
	for i, w := range words {
		ps[w[0]-'a'] = append(ps[w[0]-'a'], pair{i, 0})
	}
	for _, c := range s {
		q := ps[c-'a']
		ps[c-'a'] = nil
		for _, p := range q {
			p.j++
			if p.j == len(words[p.i]) {
				ans++
			} else {
				w := words[p.i][p.j] - 'a'
				ps[w] = append(ps[w], p)
			}
		}
	}
	return
}

func main() {
	fmt.Println(numMatchingSubseq("abcdea",
		[]string{"a", "a", "bb", "acd", "ace"}))
}
