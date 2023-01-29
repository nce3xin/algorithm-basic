package main

import "fmt"

const N int = 100003 // 取大于1e5的第一个质数，取质数冲突的概率最小

// h 开一个槽
var h [N]int

// 单链表
var e, ne [N]int
var idx int = 0

func main() {
	// init h[]
	// 将槽先清空 空指针一般用 -1 来表示
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	var n int
	fmt.Scanf("%d", &n)
	var op string
	var x int
	for ; n > 0; n-- {
		fmt.Scanf("%s%d", &op, &x)
		if op == "I" {
			insert(x)
		} else if op == "Q" {
			found := find(x)
			if found {
				fmt.Printf("Yes\n")
			} else {
				fmt.Printf("No\n")
			}
		}
	}
}

func insert(x int) {
	// 如果是负数, 那他取模也是负的, 所以 加N 再 %N 就一定是一个正数
	k := (x%N + N) % N
	// 标准的头插法，只不过头节点是h[k]
	e[idx] = x
	ne[idx] = h[k]
	h[k] = idx
	idx++
}

func find(x int) bool {
	// 用上面同样的 Hash函数，将x映射到 从 0-N 之间的数
	k := (x%N + N) % N
	for i := h[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return true
		}
	}
	return false
}
