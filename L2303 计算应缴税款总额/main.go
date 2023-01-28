package main

import "fmt"

func calculateTax1(brackets [][]int, income int) (result float64) {
	if income <= 0 {
		return result
	}
	i := 0
	for income > 0 {
		if len(brackets) == 1 {
			if income >= brackets[i][0] {
				result += float64(brackets[i][0]) * float64(brackets[i][1])
			} else {
				result += float64(income) * float64(brackets[i][1])
			}
			break
		}
		if i == 0 {
			if income >= brackets[i][0] {
				result += float64(brackets[i][0]) * float64(brackets[i][1])
			} else {
				result += float64(income) * float64(brackets[i][1])
				break
			}
		} else {
			if income >= brackets[i][0] {
				result += float64(brackets[i][0]-brackets[i-1][0]) * float64(brackets[i][1])
			} else {
				result += float64(income-brackets[i-1][0]) * float64(brackets[i][1])
				income = 0
			}
		}
		i++
		if i == len(brackets) {
			break
		}
	}
	return result / 100
}

func calculateTax(brackets [][]int, income int) float64 {
	totalTax := 0
	lower := 0
	for _, bracket := range brackets {
		upper, percent := bracket[0], bracket[1]
		tax := (min(income, upper) - lower) * percent
		totalTax += tax
		if income <= upper {
			break
		}
		lower = upper
	}
	return float64(totalTax) / 100
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	brackets := [][]int{{4, 8}, {5, 49}}
	fmt.Println(calculateTax(brackets, 2))
}
