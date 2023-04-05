package main

import "fmt"

const (
	N int = 510      // 顶点数
	M int = 1e5 + 10 // 边数
)

// 邻接表存储图
var h [N]int
var e, ne [M]int
var idx int = 0

var st [N]bool   // 每次不要重复搜同一个妹子
var match [N]int // 右边妹子对应的男生

var n1, n2, m int

func main() {
	fmt.Scanf("%d%d%d", &n1, &n2, &m)
	// init h
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	for ; m > 0; m-- {
		var u, v int
		fmt.Scanf("%d%d", &u, &v)
		// 存边只存一边就行了，虽然是无向图
		add(u, v)
	}
	res := 0
	for i := 1; i <= n1; i++ {
		clearSt() // 每次搜索前先把所有妹子清空
		if find(i) {
			res++
		}
	}
	fmt.Printf("%d\n", res)
}

func clearSt() {
	for i := 0; i < N; i++ {
		st[i] = false
	}
}

func find(x int) bool {
	for i := h[x]; i != -1; i = ne[i] {
		j := e[i] // 妹子j
		if !st[j] {
			st[j] = true
			// 如果妹子j还没匹配到男生，或者虽然匹配了男生，但我们可以给那个男生找到下家
			if match[j] == 0 || find(match[j]) {
				match[j] = x
				return true
			}
		}
	}
	return false
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
