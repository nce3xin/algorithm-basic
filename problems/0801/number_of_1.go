package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		var x int
		fmt.Scanf("%d", &x)
		res := 0
		for x != 0 {
			x -= lowBit(x)
			res++
		}
		fmt.Printf("%d ", res)
	}
}

func lowBit(x int) int {
	return x & -x
}
