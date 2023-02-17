package main

import "fmt"

func largest1BorderedSquare(grid [][]int) int {
	gx := make([][]int, len(grid)+1)
	gy := make([][]int, len(grid)+1)
	for i := 0; i <= len(grid); i++ {
		gx[i] = make([]int, len(grid[0])+1)
		gy[i] = make([]int, len(grid[0])+1)
	}
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= len(grid[i-1]); j++ {
			gx[i][j] = gx[i][j-1] + grid[i-1][j-1]
			gy[i][j] = gy[i-1][j] + grid[i-1][j-1]
		}
	}
	max := 0
	for x := 1; x < len(gx); x++ {
		for y := 1; y < len(gx[x]); y++ {
			for y1 := len(gx[x]) - 1; y1 >= y; y1-- {
				l := y1 - y + 1
				if l <= max {
					break
				}
				if x+y1-y >= len(gy) || x+y1-y >= len(gx) {
					continue
				}
				if gx[x][y1]-gx[x][y-1] != l {
					continue
				}
				if gx[x+y1-y][y1]-gx[x+y1-y][y-1] != l {
					continue
				}
				if gy[x+y1-y][y]-gy[x-1][y] != l {
					continue
				}
				if gy[x+y1-y][y1]-gy[x-1][y1] != l {
					continue
				}
				max = l
			}
		}
	}
	return max * max
}

func main() {
	fmt.Println(largest1BorderedSquare([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
}
