package main

import "fmt"

const N int = 6010

var n int
var happy [N]int

// 邻接表
var h, e, ne [N]int
var idx int = 0

// 所有状态
var f [N][2]int

// 题目没指出根节点，所以我们需要自己找根节点。根节点就是没有父节点的点，所以开一个bool数组
var hasFather [N]bool

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &happy[i])
	}
	// 初始化所有邻接表的表头
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	// 读入所有边，一共n-1条边
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scanf("%d%d", &a, &b) // b是a的父节点
		hasFather[a] = true
		add(b, a)
	}
	// 求根节点
	root := 1
	for hasFather[root] {
		root++
	}
	dfs(root)
	fmt.Printf("%d", max(f[root][0], f[root][1]))
}

func dfs(u int) {
	f[u][1] = happy[u]
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		dfs(j)

		f[u][0] += max(f[j][0], f[j][1])
		f[u][1] += f[j][0]
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
