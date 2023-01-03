package main

import "fmt"

func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	//res := fill(dp, 1.0/6.0)
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / 6.0
			}
		}
		dp = tmp
	}
	return dp
}

func fill(array []interface{}, num ...interface{}) (ans []interface{}) {
	if len(array) <= 0 {
		return
	}
	switch array[0].(type) {
	case int:
		for i := 0; i < len(array); i++ {

			ans[i] = i + 1
		}
	case int32:
		for i := 0; i < len(array); i++ {

			ans[i] = i + 1
		}
	case int64:
		fmt.Println("[]64")
	case float64:
		fmt.Println("[]float64")
	case float32:
		fmt.Println("[]float32")
	}
	fmt.Println("result", num)
	return
}

func main() {
	fmt.Println(dicesProbability(2))
}
