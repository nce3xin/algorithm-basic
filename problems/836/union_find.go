package main

import "fmt"

const N int = 1e5 + 10

var p [N]int

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)

	// init
	for i := 1; i <= n; i++ {
		p[i] = i
	}

	for ; m > 0; m-- {
		var op string
		var a, b int
		fmt.Scanf("%s", &op)
		fmt.Scanf("%d%d", &a, &b)

		if op == "M" {
			union(a, b)
		} else if op == "Q" {
			pa := find(a)
			pb := find(b)
			if pa == pb {
				fmt.Printf("Yes\n")
			} else {
				fmt.Printf("No\n")
			}
		}
	}
}

func union(a, b int) {
	p[find(a)] = find(b)
}

// 返回x的祖宗节点 + 路径压缩
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}
