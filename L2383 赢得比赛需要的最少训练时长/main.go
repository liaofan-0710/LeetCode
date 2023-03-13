package main

import "fmt"

// 两种思路
// 思路一：使用模拟
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	n := len(energy)
	time := 0
	for i := 0; i < n; i++ {
		if initialEnergy > energy[i] && initialExperience > experience[i] { // 全部超过
			initialEnergy -= energy[i]
			initialExperience += experience[i]
		} else {
			for initialEnergy < energy[i] {
				time += energy[i] - initialEnergy
				initialEnergy = energy[i]
			}
			for initialExperience < experience[i] {
				time += experience[i] - initialExperience
				initialExperience = experience[i]
			}
			if initialEnergy == energy[i] {
				time++
				initialEnergy++
			}
			if initialExperience == experience[i] {
				time++
				initialExperience++
			}
			i--
		}
	}
	return time
}

// 思路二：将所有的结果全部累加
func minNumberOfHours1(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	n := len(energy)
	time := 0
	for i := 0; i < n; i++ {
		// 精力超过是减少的
		if initialEnergy <= energy[i] {
			time += energy[i] - initialEnergy + 1
			initialEnergy = energy[i] + 1
		}
		initialEnergy -= energy[i]
		// 经验超过是增加的
		if initialExperience <= experience[i] {
			time += experience[i] - initialExperience + 1
			initialExperience = experience[i] + 1
		}
		initialExperience += experience[i]
	}
	return time
}

func main() {
	initialEnergy, initialExperience := 5, 3
	energy, experience := []int{1, 4, 3, 2}, []int{2, 6, 3, 1}
	fmt.Println(minNumberOfHours(initialEnergy, initialExperience, energy, experience))
}
