package main

func findLongestSubarray(array []string) []string {
	indices := map[int]int{}
	sum := 0
	startIndex := -1
	maxLength := 0
	indices[0] = -1
	for i, s := range array {
		if s[0] >= '0' && s[0] <= '9' {
			sum++
		} else {
			sum--
		}
		if firstIndex, ok := indices[sum]; ok {
			if i-firstIndex > maxLength {
				maxLength = i - firstIndex
				startIndex = firstIndex + 1
			}
		} else {
			indices[sum] = i
		}
	}
	if maxLength == 0 {
		return []string{}
	}
	return array[startIndex : startIndex+maxLength]
}

func main() {

}
