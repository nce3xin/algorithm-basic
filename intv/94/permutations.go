package main

import "fmt"

const N int = 10

var n int
var st [N]bool
var nums [N]int

func main() {
	fmt.Scanf("%d", &n)
	dfs(1)
}

func dfs(u int) {
	if u > n {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", nums[i])
		}
		fmt.Println()
		return
	}
	for i := 1; i <= n; i++ {
		if !st[i] {
			st[i] = true
			nums[u] = i // 当前位置设为i
			dfs(u + 1)
			st[i] = false
			nums[u] = 0
		}
	}
}
