package main

import "fmt"

var N int = 1e5 + 10

func main() {
	var n, m int
	a := make([]int, N)
	s := make([]int, N)
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
	}
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scanf("%d%d", &l, &r)
		fmt.Printf("%d\n", s[r]-s[l-1])
	}
}
