package main

import "fmt"

const N int = 1e5 + 10

var q []int = make([]int, N)
var hh int = 0
var tt int = -1

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
		} else if op == "empty" {
			isEmpty := empty()
			if isEmpty {
				fmt.Printf("YES\n")
			} else {
				fmt.Printf("NO\n")
			}
		} else if op == "query" {
			t := query()
			fmt.Printf("%d\n", t)
		} else if op == "pop" {
			pop()
		}
	}
}

func push(x int) {
	tt++
	q[tt] = x
}

func pop() {
	hh++
}

func empty() bool {
	if hh <= tt {
		return false
	}
	return true
}

func query() int {
	return q[hh]
}
