package main

import "fmt"

func reinitializePermutation(n int) (step int) {
	target := make([]int, n)
	for i := range target {
		target[i] = i
	}
	perm := append([]int(nil), target...)
	for {
		step++
		arr := make([]int, n)
		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+i/2]
			}
		}
		perm = arr
		if equal(perm, target) {
			return
		}
	}
}

func equal(a, b []int) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(reinitializePermutation(6))
}
