package main

import (
	"fmt"
	"strings"
)

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func letterCombinations(digits string) (res []string) {
	if len(digits) == 0 {
		return res
	} else if len(digits) == 1 {
		return strings.Split(phoneMap[string(digits[0])], "")
	} else if len(digits) == 2 {
		nums1 := strings.Split(phoneMap[string(digits[0])], "")
		nums2 := strings.Split(phoneMap[string(digits[1])], "")
		for i := 0; i < len(nums1); i++ {
			for j := 0; j < len(nums2); j++ {
				res = append(res, nums1[i]+nums2[j])
			}
		}
	} else if len(digits) == 3 {
		nums1 := strings.Split(phoneMap[string(digits[0])], "")
		nums2 := strings.Split(phoneMap[string(digits[1])], "")
		nums3 := strings.Split(phoneMap[string(digits[2])], "")
		for i := 0; i < len(nums1); i++ {
			for j := 0; j < len(nums2); j++ {
				for k := 0; k < len(nums3); k++ {
					res = append(res, nums1[i]+nums2[j]+nums3[k])
				}
			}
		}
	} else if len(digits) == 4 {
		nums1 := strings.Split(phoneMap[string(digits[0])], "")
		nums2 := strings.Split(phoneMap[string(digits[1])], "")
		nums3 := strings.Split(phoneMap[string(digits[2])], "")
		nums4 := strings.Split(phoneMap[string(digits[3])], "")
		for i := 0; i < len(nums1); i++ {
			for j := 0; j < len(nums2); j++ {
				for k := 0; k < len(nums3); k++ {
					for s := 0; s < len(nums4); s++ {
						res = append(res, nums1[i]+nums2[j]+nums3[k]+nums4[s])
					}
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(letterCombinations("23"))
}
