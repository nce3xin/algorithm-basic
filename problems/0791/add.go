package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	var A, B []int
	for i := len(a) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(a[i]))
		A = append(A, digit)
	}
	for i := len(b) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(b[i]))
		B = append(B, digit)
	}

	C := add(A, B)
	
	for i := len(C) - 1; i >= 0; i-- {
		fmt.Printf("%d", C[i])
	}
}

func add(A []int, B []int) []int {
	var C []int
	t := 0 // 进位
	for i := 0; i < len(A) || i < len(B); i++ {
		if i < len(A) {
			t += A[i]
		}
		if i < len(B) {
			t += B[i]
		}
		C = append(C, t%10)
		t /= 10
	}
	if t > 0 {
		C = append(C, t)
	}
	return C
}
