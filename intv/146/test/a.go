package main

import (
	"fmt"
	"sort"
)

const N int = 15

var a [N]int

func main() {
	n := 10
	for i := 0; i < n; i++ {
		a[i] = n - i
	}
	fmt.Printf("before sorted:\n")
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	sort.Ints(a[:n])
	fmt.Printf("after sorted:\n")
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
