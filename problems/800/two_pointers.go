package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 10

var a []int = make([]int, N)
var b []int = make([]int, N)

func main() {
	var n, m, x int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m, &x)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	for i, j := 0, m-1; i < n; i++ {
		for j >= 0 && a[i]+b[j] > x {
			j--
		}
		if a[i]+b[j] == x {
			fmt.Printf("%d %d", i, j)
			break
		}
	}
}
