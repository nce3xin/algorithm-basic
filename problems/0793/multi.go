package main

import (
	"fmt"
	"strconv"
)

func main() {
	// input
	var a string
	var b int // b是一个很小的数，直接用int读入
	fmt.Scanf("%s", &a)
	fmt.Scanf("%d", &b)
	var A []int
	for i := len(a) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(a[i]))
		A = append(A, c)
	}

	C := mul(A, b)

	for i := len(C) - 1; i >= 0; i-- {
		fmt.Printf("%d", C[i])
	}
}

func mul(A []int, b int) []int {
	var C []int
	t := 0 // 进位
	// 把最后有进位的情况合并，t不是0的话，就一直循环
	for i := 0; i < len(A) || t != 0; i++ {
		if i < len(A) {
			t += A[i] * b
		}
		C = append(C, t%10)
		t /= 10
	}
	// 比如123*0等于000，要去除前导0
	for len(C) > 1 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}
	return C
}
