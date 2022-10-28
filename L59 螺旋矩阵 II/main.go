package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		nums := make([]int, n)
		res = append(res, nums)
	}
	t, r, d, l := 0, n-1, n-1, 0
	num := 1
	for num <= n*n {
		// 上  从左到右
		for i := l; i <= r; i++ {
			res[t][i] = num
			num++
		}
		t++
		// 右 从上到下
		for i := t; i <= d; i++ {
			res[i][r] = num
			num++
		}
		r--
		// 下  从右到左
		for i := r; i >= l; i-- {
			res[d][i] = num
			num++
		}
		d--
		// 左 从下到上
		for i := d; i >= t; i-- {
			res[i][l] = num
			num++
		}
		l++
	}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
