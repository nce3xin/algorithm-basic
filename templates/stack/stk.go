package stack

// 数组模拟栈

const N int = 1e5 + 10

// 用top表示栈顶所在的索引。初始时，top = -1。表示没有元素。
var top int = -1
var stk []int = make([]int, N)

// push 栈顶所在索引往后移动一格，然后放入x
func push(x int) {
	top++
	stk[top] = x
}

// top 往前移动一格
func pop() {
	top--
}

// top 大于等于 0 栈非空，小于 0 栈空
func empty() bool {
	if top >= 0 {
		return false
	}
	return true
}

// 返回栈顶元素
func getTop() int {
	x := stk[top]
	return x
}
