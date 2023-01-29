package main

import "fmt"

const N int = 10

var n int
var path [N]int // 用 path 数组保存排列
var vis [N]bool // 用 vis 数组表示数字是否用过

func main() {
	fmt.Scanf("%d", &n)
	// 填第0个位置的数
	// dfs(i) 表示的含义是：在 path[i] 处填写数字，然后递归的在下一个位置填写数字。
	dfs(0)
}

func dfs(u int) {
	// 用 path 数组保存排列，当排列的长度为 n 时，是一种方案，输出。
	if u == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Printf("\n")
		return
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			path[u] = i
			vis[i] = true
			dfs(u + 1) // 走到下一层，填第u+1个位置的数
			// 当dfs函数结束的时候，意味着下面所有的路都已经走完了，就要回溯了。回来的时候要恢复现场。
			// path[u] = 0 这句其实可以不写，因为下一次path[u]上的数会被不断覆盖掉
			vis[i] = false
		}
	}
}
