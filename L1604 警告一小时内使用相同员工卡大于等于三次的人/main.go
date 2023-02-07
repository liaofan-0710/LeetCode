package main

import (
	"fmt"
	"sort"
)

func alertNames(keyName []string, keyTime []string) []string {
	var res []string
	timeMap := map[string][]int{}
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		timeMap[name] = append(timeMap[name], hour*60+minute)
	}
	for name, times := range timeMap {
		sort.Ints(times)
		for i, t := range times[2:] {
			if t-times[i] <= 60 {
				res = append(res, name)
				break
			}
		}
	}
	sort.Strings(res)
	return res
}

func main() {
	str := []string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b"}
	keyTime := []string{"23:27", "03:14", "12:57", "13:35", "13:18", "21:58", "22:39", "10:49", "19:37", "14:14", "10:41"}
	fmt.Println(alertNames(str, keyTime))
}
