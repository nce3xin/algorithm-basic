package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 10

var n, m int
var a, b [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}
	i, j := 0, 0
	for i < n && j < m {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	if i == n {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
