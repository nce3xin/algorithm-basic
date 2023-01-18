package main

import "fmt"

// N 数的个数
const N int = 1e5 + 10

// M Trie树中节点个数最大是多少呢？最多有N个数，每个数有31位，所以节点个数最多是31*N
const M int = 31 * N

var a [N]int      // read input
var son [M][2]int // Trie树存储
var idx int = 0   // 当前用到哪个下标了

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	res := 0

	// add all int to trie
	for i := 0; i < n; i++ {
		insert(a[i])
	}

	for i := 0; i < n; i++ {
		//insert(a[i])
		t := query(a[i])
		xor := a[i] ^ t
		res = max(res, xor)
	}
	fmt.Printf("%d", res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func insert(x int) {
	// 从根节点开始
	p := 0
	for i := 30; i >= 0; i-- {
		// 取 x 的第i位的二进制数是什么  x >> k & 1 (前面的模板)
		u := x >> i & 1
		if son[p][u] == 0 { // 子节点不存在
			idx++
			son[p][u] = idx
		}
		p = son[p][u] // 指针指向下一层
	}
}

// 返回与x异或最大的数
func query(x int) int {
	// 从根节点开始
	p := 0
	res := 0
	// 从最大位开始找
	for i := 30; i >= 0; i-- {
		u := x >> i & 1
		if son[p][reverse(u)] > 0 { // 相反的分支存在
			p = son[p][reverse(u)]
			// 求这条路径上的数，指针每下移一层，res就左移一位，再加上当前bit
			res = res<<1 + reverse(u)
		} else {
			p = son[p][u]
			res = res<<1 + u
		}
	}
	return res
}

func reverse(x int) int {
	res := 0
	if x == 1 {
		res = 0
	} else if x == 0 {
		res = 1
	}
	return res
}
