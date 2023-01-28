package main

import (
	"fmt"
	"strconv"
)

func rearrangeCharacters(s, target string) int {
	var cntS, cntT [26]int
	for _, c := range s {
		cntS[c-'a']++
	}
	for _, c := range target {
		cntT[c-'a']++
	}
	ans := len(s)
	for i, c := range cntT {
		if c > 0 {
			ans = min(ans, cntS[i]/c)
			if ans == 0 {
				return 0
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func differenceOfSum(nums []int) int {
	sum1, sum2 := 0, 0
	for i := 0; i < len(nums); i++ {
		sum1 += nums[i]
		str := strconv.Itoa(nums[i])
		num := 0
		for j := 0; j < len(str); j++ {
			atoi, _ := strconv.Atoi(string(str[j]))
			num += atoi
		}
		sum2 += num
	}
	if sum1-sum2 > 0 {
		return sum1 - sum2
	}
	return sum2 - sum1
}

func main() {
	fmt.Println(rearrangeCharacters("ilovecodingonleetcode", "code"))
}
