package main

import (
	"fmt"
	"sort"
)

func vowelStrings(words []string, left int, right int) (ans int) {
	count := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}
	for i := left; i <= right; i++ {
		_, ok := count[words[i][0]]
		_, ok1 := count[words[i][len(words[i])-1]]
		if ok && ok1 {
			ans++
		}
	}
	return ans
}

func maxScore(nums []int) (ans int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	prefix := make([]int, len(nums))
	if len(prefix) > 0 {
		prefix[0] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	for i := 0; i < len(prefix); i++ {
		if prefix[i] > 0 {
			ans++
		}
	}
	return
}

func beautifulSubarrays(nums []int) (ans int64) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] ^ x
	}
	cnt := map[int]int{}
	for _, x := range s {
		// 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
		ans += int64(cnt[x])
		cnt[x]++
	}
	return
}

// 优化：合并
func beautifulSubarrays1(nums []int) (ans int64) {
	s := 0
	cnt := map[int]int{s: 1} // s[0]
	for _, x := range nums {
		s ^= x
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}

func main() {
	//fmt.Println(maxScore([]int{-687767, -860350, 950296, 52109, 510127, 545329, -291223, -966728, 852403, 828596, 456730, -483632, -529386, 356766, 147293, 572374, 243605, 931468, 641668, 494446}))
	fmt.Println(beautifulSubarrays([]int{4, 3, 1, 2, 4}))
}
