package main

import "fmt"

// 第一种 深度优先搜索
func possibleBipartition1(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	color := make([]int, n) // color[x] = 0 表示未访问节点 x
	var dfs func(int, int) bool
	dfs = func(x, c int) bool {
		color[x] = c
		for _, y := range g[x] {
			if color[y] == c || color[y] == 0 && !dfs(y, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

// 第二种 广度优先搜索
func possibleBipartition2(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	color := make([]int, n) // 0 表示未访问该节点
	for i, c := range color {
		if c == 0 {
			q := []int{i}
			color[i] = 1
			for len(q) > 0 {
				x := q[0]
				q = q[1:]
				for _, y := range g[x] {
					if color[y] == color[x] {
						return false
					}
					if color[y] == 0 {
						color[y] = -color[x]
						q = append(q, y)
					}
				}
			}
		}
	}
	return true
}

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return unionFind{parent, make([]int, n)}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

func (uf unionFind) isConnected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

// 第三种 并查集
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	uf := newUnionFind(n)
	for x, nodes := range g {
		for _, y := range nodes {
			uf.union(nodes[0], y)
			if uf.isConnected(x, y) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
}
