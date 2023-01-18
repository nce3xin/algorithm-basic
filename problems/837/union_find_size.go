package main

import "fmt"

const N int = 1e5 + 10

var p [N]int
var size [N]int

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	// init
	for i := 1; i <= n; i++ {
		p[i] = i
		size[i] = 1
	}
	for ; m > 0; m-- {
		var op string
		fmt.Scanf("%s", &op)
		if op == "C" {
			var a, b int
			fmt.Scanf("%d%d", &a, &b)
			if find(a) == find(b) {
				continue
			}
			union(a, b)
		} else if op == "Q1" {
			var a, b int
			fmt.Scanf("%d%d", &a, &b)
			if find(a) == find(b) {
				fmt.Printf("Yes\n")
			} else {
				fmt.Printf("No\n")
			}
		} else if op == "Q2" {
			var a int
			fmt.Scanf("%d", &a)
			sz := size[find(a)]
			fmt.Printf("%d\n", sz)
		}
	}
}

func union(a, b int) {
	// 这两句顺序不能变
	size[find(b)] += size[find(a)]
	p[find(a)] = find(b)
}

// 返回x的祖宗节点
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}
