package main

import "fmt"

func threeEqualParts1(arr []int) []int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}

	partial := sum / 3
	first, second, third, cur := 0, 0, 0, 0
	for i, x := range arr {
		if x == 1 {
			if cur == 0 {
				first = i
			} else if cur == partial {
				second = i
			} else if cur == 2*partial {
				third = i
			}
			cur++
		}
	}

	n := len(arr)
	length := n - third
	if first+length <= second && second+length <= third {
		i := 0
		for third+i < n {
			if arr[first+i] != arr[second+i] || arr[first+i] != arr[third+i] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + length - 1, second + length}
	}
	return []int{-1, -1}
}

func threeEqualParts(arr []int) []int {
	find := func(x int) int {
		s := 0
		for i, v := range arr {
			s += v
			if s == x {
				return i
			}
		}
		return 0
	}
	n := len(arr)
	cnt := 0
	for _, v := range arr {
		cnt += v
	}
	if cnt%3 != 0 {
		return []int{-1, -1}
	}
	if cnt == 0 {
		return []int{0, n - 1}
	}
	cnt /= 3
	i, j, k := find(1), find(cnt+1), find(cnt*2+1)
	for ; k < n && arr[i] == arr[j] && arr[j] == arr[k]; i, j, k = i+1, j+1, k+1 {
	}
	if k == n {
		return []int{i - 1, j}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(threeEqualParts([]int{1, 1, 0, 1, 1}))
}
