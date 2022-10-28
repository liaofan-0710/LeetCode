package main

import "fmt"

func DigitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number = number / 10
	}
	return sum
}

func CanReach(boad [][]bool, m, n, i, j int) bool {
	return boad[i][j] ||
		(j-1 >= 0 && boad[i][j-1]) ||
		(i-1 >= 0 && boad[i-1][j]) ||
		(j+1 < n && boad[i][j+1]) ||
		(i+1 < m && boad[i+1][j])
}

func movingCount(m int, n int, k int) int {
	ans := 0
	boad := make([][]bool, m)
	for i := 0; i < m; i++ {
		boad[i] = make([]bool, n)
	}

	boad[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if DigitSum(i)+DigitSum(j) <= k && CanReach(boad, m, n, i, j) {
				boad[i][j] = true
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(movingCount(2, 3, 1))
}
