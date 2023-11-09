package main

import "fmt"

const N int = 20

var n int
var st [N]bool

func main() {
	fmt.Scanf("%d", &n)
	dfs(1)
}

func dfs(u int) {
	if u > n {
		for i := 1; i <= n; i++ {
			if st[i] {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
		return // 一定不要忘了return!
	}
	st[u] = true
	dfs(u + 1)

	st[u] = false
	dfs(u + 1)
}
