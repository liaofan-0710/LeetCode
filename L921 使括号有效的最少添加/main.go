package main

import "fmt"

func minAddToMakeValid1(s string) int {
	nums := []byte{}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			nums = append(nums, s[i])
		} else {
			if len(nums) > 0 && nums[len(nums)-1] == '(' {
				nums = nums[:len(nums)-1]
			} else {
				result++
			}
		}
	}
	return result + len(nums)
}

func minAddToMakeValid(s string) (ans int) {
	cnt := 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			ans++
		}
	}
	return ans + cnt
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
}
