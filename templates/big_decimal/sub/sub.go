package sub

import (
	"fmt"
	"strconv"
)

func Demo() {
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
	if cmp(A, B) { // 如果A >= B
		C := sub(A, B)

		for i := len(C) - 1; i >= 0; i-- {
			fmt.Printf("%d", C[i])
		}
	} else {
		C := sub(B, A)

		fmt.Printf("-")
		for i := len(C) - 1; i >= 0; i-- {
			fmt.Printf("%d", C[i])
		}
	}
}

// A>=B, return true
func cmp(A, B []int) bool {
	if len(A) != len(B) {
		return len(A) > len(B)
	}
	// 从高位开始比, 高位存在数组的后面
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] != B[i] {
			return A[i] > B[i]
		}
	}
	return true
}

// 模板要求必须满足A >= B
// 如果A < B，就计算B-A，结果前面再加个负号。
func sub(A, B []int) []int {
	var C []int // 结果数组
	t := 0      // 借位
	// 能以A的长度为准，是因为该函数要求 A > B
	// 从整数的个位开始减，个位保存在数组的最前面
	for i := 0; i < len(A); i++ {
		// C[i] = A[i] - B[i] - t
		// 用进位t存储当前位的减法结果
		// 不能直接减B[i], 因为B[i]不一定存在，要判断
		t = A[i] - t
		if i < len(B) {
			t -= B[i]
		}
		// 用进位t存储当前位的减法结果
		// t>=0, t = t
		// t < 0, t = t + 10
		// (t+10)%10 可以统一这两种情况
		C = append(C, (t+10)%10)
		// 如果当前位的计算结果t是小于0的，说明需要借位，令t=1
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	// 比如113-110=3，上面的代码结果是003(数组中倒着存，是300)，需要去掉前导0
	// len(C) > 1 是因为，例如12-12，模板计算结果是00，需要保留一个0，而不是把所有0都删除
	for len(C) > 1 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}
	return C
}
