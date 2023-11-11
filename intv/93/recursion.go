package main

import "fmt"

const N int = 30

var n, m int
var st [N]bool

func main() {
	fmt.Scanf("%d%d", &n, &m)
	dfs(1, 0)
}

// s表示选的数的个数
func dfs(u int, s int) {
	if s == m { // 够m个数了，后面的都可以不要了，输出
		for i := 1; i <= n; i++ {
			if st[i] {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
		return
	}
	// 搜到最后都不够m个数，直接返回
	if u > n {
		return
	}
	// 2种情况，选或不选
	// 必须要把选的情况放在前面，因为题目要求字典序输出
	st[u] = true
	dfs(u+1, s+1)
	st[u] = false // 恢复现场

	dfs(u+1, s)
}
