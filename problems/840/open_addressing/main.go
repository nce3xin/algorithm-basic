package main

import "fmt"

// 开放寻址法一般开 数据范围的 2~3倍, 这样大概率就没有冲突了

const N int = 200003        // 大于数据范围的第一个质数
const null int = 0x3f3f3f3f // 规定空指针为 null 0x3f3f3f3f

var h [N]int

func find(x int) int {
	k := (x%N + N) % N
	for h[k] != null && h[k] != x {
		k++
		if k == N {
			k = 0
		}
	}
	return k // 如果这个位置是空的, 则返回的是他应该存储的位置
}

func main() {
	// init h
	for i := 0; i < N; i++ {
		h[i] = null // 规定空指针为 0x3f3f3f3f
	}

	var n int
	fmt.Scanf("%d", &n)
	var op string
	var x int
	for ; n > 0; n-- {
		fmt.Scanf("%s%d", &op, &x)
		if op == "I" {
			k := find(x)
			h[k] = x
		} else if op == "Q" {
			k := find(x)
			if h[k] != null {
				fmt.Printf("Yes\n")
			} else {
				fmt.Printf("No\n")
			}
		}
	}
}
