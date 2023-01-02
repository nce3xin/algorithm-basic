package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var b int
	fmt.Scanf("%s", &a)
	fmt.Scanf("%d", &b)
	var A []int
	for i := len(a) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(a[i]))
		A = append(A, c)
	}

	C, r := div(A, b)

	for i := len(C) - 1; i >= 0; i-- {
		fmt.Printf("%d", C[i])
	}
	fmt.Printf("\n%d", r)
}

func div(A []int, b int) ([]int, int) {
	var C []int // 商
	r := 0      // 余数
	// 从整数的最高位开始，和加减乘不一样
	for i := len(A) - 1; i >= 0; i-- {
		r = r*10 + A[i]
		C = append(C, r/b)
		r %= b
	}
	reverse(C)
	// 去除前导0, 例如 1000/9,上面计算的商是0111，前导0应该去掉
	for len(C) > 1 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}
	return C, r
}

func reverse(A []int) {
	var C []int
	for i := len(A) - 1; i >= 0; i-- {
		C = append(C, A[i])
	}
	copy(A, C)
}
