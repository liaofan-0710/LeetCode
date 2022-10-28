package main

import (
	"fmt"
	"strings"
)

func reformatNumber1(number string) string {
	nums := strings.Split(number, " ")
	str := ""
	for i := 0; i < len(nums); i++ {
		str += nums[i]
	}
	nums = strings.Split(str, "-")
	str = ""
	for i := 0; i < len(nums); i++ {
		str += nums[i]
	}
	nums = strings.Split(str, "")
	str = ""
	count, boolean := 0, true
	if len(nums) == 4 {
		for i := 0; i < len(nums); i++ {
			str += nums[i]
			count++
			if count == 2 && i != len(nums)-1 {
				str += "-"
				count = 0
			}
		}
		return str
	}
	for i := 0; i < len(nums); i++ {
		str += nums[i]
		count++
		if count == 3 && boolean && i != len(nums)-1 {
			str += "-"
			count = 0
			if len(nums)-i == 5 {
				boolean = false
			}
		} else if count == 2 && boolean == false && i != len(nums)-1 {
			str += "-"
			count = 0
		}
	}
	return str
}

func reformatNumber2(number string) string {
	s := strings.ReplaceAll(number, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	ans := []string{}
	i := 0
	for ; i+4 < len(s); i += 3 {
		ans = append(ans, s[i:i+3])
	}
	s = s[i:]
	if len(s) < 4 {
		ans = append(ans, s)
	} else {
		ans = append(ans, s[:2], s[2:])
	}
	return strings.Join(ans, "-")
}

func reformatNumber(number string) string {
	nums := strings.Split(number, " ")
	str := ""
	for i := 0; i < len(nums); i++ {
		str += nums[i]
	}
	nums = strings.Split(str, "-")
	str = ""
	for i := 0; i < len(nums); i++ {
		str += nums[i]
	}
	ans := []string{}
	i := 0
	for ; i+4 < len(str); i += 3 {
		ans = append(ans, str[i:i+3])
	}
	str = str[i:]
	if len(str) < 4 {
		ans = append(ans, str)
	} else {
		ans = append(ans, str[:2], str[2:])
	}
	return strings.Join(ans, "-")
}

func main() {
	fmt.Println(reformatNumber("9964-"))
}
