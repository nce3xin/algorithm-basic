package main

const N int = 1010

var stk [N]int
var tt int = 0

func isPopOrder(pushV []int, popV []int) bool {
	if len(pushV) != len(popV) {
		return false
	}
	i := 0 // 遍历出栈序列
	for _, item := range pushV {
		push(item) // 依次喂数
		// 检查当前出栈的数是否等于栈顶元素
		for !empty() && top() == popV[i] {
			pop()
			i++
		}
	}
	if empty() {
		return true
	}
	return false
}

func push(x int) {
	tt++
	stk[tt] = x
}

func pop() {
	tt--
}

func top() int {
	return stk[tt]
}

func empty() bool {
	if tt > 0 {
		return false
	}
	return true
}
