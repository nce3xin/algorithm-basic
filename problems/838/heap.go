package main

import "fmt"

const N int = 1e5 + 10

// h[N]存储堆中的值, 下标从1开始，h[1]是堆顶，x的左儿子是2x, 右儿子是2x + 1
var h [N]int
var size int // 堆中元素个数

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	size = n
	// 下标从1开始
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &h[i])
	}
	// O(N) 建堆
	for i := n / 2; i > 0; i-- {
		down(i)
	}
	for ; m > 0; m-- {
		fmt.Printf("%d ", h[1])
		h[1] = h[size]
		size--
		down(1)
	}
}

// 小根堆
// u is index
func down(u int) {
	// t 表示 三个点里的最小值。三个点分别是u, u的左儿子, u的右儿子
	t := u
	// u的左儿子存在且小于u
	if 2*u <= size && h[2*u] < h[t] {
		t = 2 * u
	}
	// u的右儿子存在且小于u
	if 2*u+1 <= size && h[2*u+1] < h[t] {
		t = 2*u + 1
	}
	if t != u { // 说明根节点不是最小的
		// swap
		h[t], h[u] = h[u], h[t]
		// 递归处理
		down(t)
	}
}

// 这题其实并不需要up操作，这里写出up操作，只是为了给出堆的完整实现
// up比down简单
func up(u int) {
	// u/2表示根节点，u/2 > 0 根节点存在
	for u/2 > 0 && h[u/2] > h[u] {
		// swap
		h[u/2], h[u] = h[u], h[u/2]
		u = u / 2
	}
}
