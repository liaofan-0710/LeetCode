package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits1(cpdomains []string) []string {
	countMap := make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		sum := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(sum[0])
		str := strings.Split(sum[1], ".")
		s := str[len(str)-1]
		for j := len(str) - 2; j >= -1; j-- {
			if _, ok := countMap[s]; !ok {
				countMap[s] = count
			} else {
				countMap[s] += count
			}
			if j <= -1 {
				break
			}
			s = str[j] + "." + s
		}
	}
	result := []string{}
	for k, v := range countMap {
		result = append(result, strconv.Itoa(v)+" "+k)
	}
	return result
}

func subdomainVisits(cpdomains []string) []string {
	cnt := map[string]int{}
	for _, s := range cpdomains {
		i := strings.IndexByte(s, ' ')
		v, _ := strconv.Atoi(s[:i])
		for ; i < len(s); i++ {
			if s[i] == ' ' || s[i] == '.' {
				cnt[s[i+1:]] += v
			}
		}
	}
	ans := make([]string, 0, len(cnt))
	for s, v := range cnt {
		ans = append(ans, strconv.Itoa(v)+" "+s)
	}
	return ans
}

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
