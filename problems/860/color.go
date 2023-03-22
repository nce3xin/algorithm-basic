package main

import "fmt"

var n, m int

const N int = 1e5 + 10 // 点的个数
const M int = 2e5 + 10 // 边的个数，因为是无向图，所以边的数量是点的数量的两倍
// 邻接表存储稀疏图
var h [N]int
var e, ne [M]int
var idx int = 0

// 存储点的颜色
var color [N]int // 两种颜色，1和2

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	// init h
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	for ; m > 0; m-- {
		var a, b int
		fmt.Scanf("%d%d", &a, &b)
		add(a, b)
		add(b, a)
	}
	flag := true // 标记图是否是二分图，true表示是，false表示不是
	// 开始染色过程
	for i := 1; i <= n; i++ {
		if color[i] == 0 { // 如果第i个点还没有被染色
			// 把第i个点所在的连通块都染一遍
			if !dfs(i, 1) { // 如果有冲突发生
				flag = false
				break
			}
		}
	}

	if flag {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}

// dfs 返回false表示染色过程中有冲突发生
// dfs 参数c表示要染成什么颜色
func dfs(u int, c int) bool {
	color[u] = c
	// 遍历点u的所有临边
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if color[j] == 0 { // 如果相邻点j没有被染色, 注意这里是等于0表示没有被染色，别写反了
			// 3-c正好能描述：当u染1时，j染2。u染2时，j染1
			if !dfs(j, 3-c) { // 如果染色过程中有冲突
				return false
			}
		} else if color[j] == c { // 相连点u和j居然染了相同颜色，冲突！
			return false
		}
	}
	return true
}
