package main

import (
	"fmt"
)

const N int = 1e5 + 10

var s string
var k int
var n int

// 用栈存字符串
var stk [N]byte
var tt int = 0

func main() {
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &k)
	n = len(s)
	for i := 0; i < n; i++ {
		// 找到有逆序对，就循环删除栈顶元素
		for k > 0 && tt > 0 && stk[tt] > s[i] {
			tt--
			k--
		}
		// 入栈
		tt++
		stk[tt] = s[i]
	}
	// 如果栈里已经是单调递增序列了，但是k还没用完，就把后缀删掉
	for k != 0 {
		tt--
		k--
	}
	// 找第一个不是前导0的位置
	i := 1 // 因为stk栈是从下标1开始存数的
	// 这里一定是单引号括起来的'0', 因为stk里的数据类型是byte
	for i <= tt && stk[i] == '0' {
		i++
	}
	// 如果所有数都是0，就直接输出0
	if i > tt {
		fmt.Printf("0")
		return
	}
	// 从不为前导0的位置i开始输出
	for j := i; j <= tt; j++ {
		fmt.Printf("%d", stk[j]-'0') // 这里stk[j]一定要记得减'0',否则会输出好长一段数
	}
}
