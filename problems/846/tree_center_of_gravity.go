package main

import "fmt"

const N int = 1e5 + 10 // 点的数量
const M int = 2 * N    // 边的数量，最多是点个数的两倍，因为是无向图

var n int

// graph storage
var h [N]int // 每个点对应的链表的头节点的指针
var e [M]int
var ne [M]int
var idx int = 0
var ans int = N // 最终结果
var vis [N]bool // 记录点是否已经遍历过

func main() {
	// 注意，init h[]一定要在add()前面，不然add的边就不对了，我就因为这个一直没AC
	// init h[] to -1
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	fmt.Scanf("%d", &n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scanf("%d%d", &a, &b)
		// 因为是无向图，所以两个方向的边都要添加
		add(a, b)
		add(b, a)
	}

	dfs(1)
	fmt.Printf("%d", ans)
}

// dfs 返回以u为根的子树的点的数量
func dfs(u int) int {
	vis[u] = true // 标记一下，已经被搜过了
	res := 0      // 删除节点u以后，剩余各个连通块中，点数的最大值
	sum := 1      // 以u为根的子树的点的数量
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]    // 链表上节点的值，也就是节点编号，题目中的节点编号是从1开始的
		if !vis[j] { // 如果j节点没有被访问过
			s := dfs(j)
			res = max(res, s)
			sum += s // 因为以j为根的子树也是u的一部分
		}
	}
	res = max(res, n-sum)
	ans = min(ans, res)
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 添加一条边a->b
func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
