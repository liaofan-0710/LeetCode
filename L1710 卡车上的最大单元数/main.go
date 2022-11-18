package main

import (
	"fmt"
	"sort"
)

func maximumUnits1(boxTypes [][]int, truckSize int) (res int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for i := 0; i < len(boxTypes); i++ {
		for j := 0; j < boxTypes[i][0]; j++ {
			if truckSize > 0 {
				res += boxTypes[i][1]
				truckSize--
			} else {
				break
			}
		}
	}
	return
}

func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	for _, p := range boxTypes {
		if p[0] >= truckSize {
			ans += truckSize * p[1]
			break
		}
		truckSize -= p[0]
		ans += p[0] * p[1]
	}
	return
}

func main() {
	boxTypes := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	truckSize := 10
	fmt.Println(maximumUnits(boxTypes, truckSize))
}
