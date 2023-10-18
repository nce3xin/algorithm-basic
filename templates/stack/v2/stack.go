package v2

// 数组模拟栈

const N int = 1e5 + 10

// tt表示栈顶，tt=0这个位置不存数据
var tt int = 0
var stk []int = make([]int, N)

// push 向栈顶插入一个数
func push(x int) {
	tt++
	stk[tt] = x
}

// top 从栈顶弹出一个数
func pop() {
	tt--
}

// tt 大于 0 栈非空，小于等于 0 栈空
func empty() bool {
	if tt > 0 {
		return false
	}
	return true
}

// 返回栈顶元素
func getTop() int {
	x := stk[tt]
	return x
}
