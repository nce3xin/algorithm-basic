package main

import "fmt"

const N int = 1e5 + 10

// 用top表示栈顶所在的索引。初始时，top = -1。表示没有元素。
var top int = -1
var stk []int = make([]int, N)

func main() {
	var M int
	fmt.Scanf("%d", &M)
	var op string
	for ; M > 0; M-- {
		fmt.Scanf("%s", &op)
		if op == "push" {
			var x int
			fmt.Scanf("%d", &x)
			push(x)
		} else if op == "query" {
			t := getTop()
			fmt.Printf("%d\n", t)
		} else if op == "pop" {
			pop()
		} else if op == "empty" {
			isEmpty := empty()
			if isEmpty {
				fmt.Printf("YES\n")
			} else {
				fmt.Printf("NO\n")
			}
		}
	}
}

func push(x int) {
	top++
	stk[top] = x
}

func pop() {
	top--
}

func empty() bool {
	if top >= 0 {
		return false
	}
	return true
}

func getTop() int {
	x := stk[top]
	return x
}
