package main

import "fmt"

func validateStackSequences1(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	res := make([]int, 0)
	j, i, s := 0, 0, 0
	for i < len(pushed) || s < len(popped) {
		if j == len(popped) {
			break
		}
		for k := len(res) - 1; len(res) > 0 && res[k] == popped[j]; k-- {
			res = res[:len(res)-1]
			j++
			s++
		}
		if i < len(pushed) && pushed[i] != popped[j] {
			res = append(res, pushed[i])
		} else if i < len(pushed) && pushed[i] == popped[j] {
			j++
			s++
		} else {
			s++
		}
		i++
	}
	return len(res) == 0
}

func validateStackSequences2(pushed []int, popped []int) bool {
	var stack [1001]int
	j1 := 0
	j2 := 0
	stack[0] = -1
	for i, _ := range pushed {
		if j1 < 0 {
			j1 = 0
		}
		stack[j1] = pushed[i]
		j1++
		for j1 > 0 && stack[j1-1] == popped[j2] {
			stack[j1-1] = -1
			j1--
			j2++
		}
	}
	return stack[0] == -1
}

// 大道至简
func validateStackSequences(pushed []int, popped []int) bool {
	i, j := 0, 0
	for _, v := range pushed {
		pushed[i] = v
		for i >= 0 && pushed[i] == popped[j] {
			j++
			i--
		}
		i++
	}
	return j == len(popped)
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
