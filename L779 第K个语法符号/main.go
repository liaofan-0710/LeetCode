package main

import (
	"fmt"
)

func kthGrammar1(n, k int) int {
	if n == 1 {
		return 0
	}
	return k&1 ^ 1 ^ kthGrammar(n-1, (k+1)/2)
}

func kthGrammar2(n, k int) int {
	if k == 1 {
		return 0
	}
	if k > 1<<(n-2) {
		return 1 ^ kthGrammar(n-1, k-1<<(n-2))
	}
	return kthGrammar(n-1, k)
}

func kthGrammar(n, k int) (ans int) {
	//return bits.OnesCount(uint(k-1)) & 1
	for k--; k > 0; k &= k - 1 {
		ans ^= 1
	}
	return
}

func main() {
	fmt.Println(kthGrammar(3, 2))
}
