package main

import "fmt"

const N int = 1e5 + 10

// 双栈, 一个存数字，一个存运算符
var num = NewStack(N)
var op = NewStack(N)

type Stack struct {
	stk []interface{}
	tt  int
}

func NewStack(N int) *Stack {
	return &Stack{
		stk: make([]interface{}, N),
		tt:  0,
	}
}

func (s *Stack) Push(x interface{}) {
	s.tt++
	s.stk[s.tt] = x
}

func (s *Stack) Pop() {
	s.tt--
}

func (s *Stack) Top() interface{} {
	return s.stk[s.tt]
}

func (s *Stack) Empty() bool {
	if s.tt > 0 {
		return false
	}
	return true
}

func main() {
	// 定义运算符优先级，加减法的运算符优先级是1，乘除法的运算符优先级是2。数字越大，优先级越高
	pr := make(map[rune]int)
	pr['+'] = 1
	pr['-'] = 1
	pr['*'] = 2
	pr['/'] = 2

	var s string // input
	fmt.Scanf("%s", &s)
	str := []rune(s) // convert to []rune
	for i := 0; i < len(str); i++ {
		c := str[i]
		// 如果是数字，把数字扣出来
		if isDigit(c) {
			x := 0
			j := i
			for j < len(s) && isDigit(str[j]) {
				x = 10*x + int(s[j]-'0')
				j++
			}
			// 不是i=j, 因为外层for循环中，i还会再加1
			i = j - 1
			num.Push(x)
		} else if c == '(' {
			op.Push(c)
		} else if c == ')' { // 遇到右括号计算括号里面的
			for op.Top() != '(' { // 括号里从右到左计算，一直计算到左括号
				eval()
			}
			op.Pop() // 左括号出栈
		} else {
			// 当前运算符比之前运算符的优先级低，则先计算
			for !op.Empty() && op.Top() != '(' && pr[op.Top().(rune)] >= pr[c] {
				eval() // 操作一次
			}
			op.Push(c)
		}
	}
	for !op.Empty() { // 计算剩余的
		eval() // 操作一次
	}
	fmt.Printf("%d", num.Top())
}

// 求值
func eval() {
	b := num.Top().(int)
	num.Pop()
	a := num.Top().(int)
	num.Pop()
	c := op.Top().(rune)
	op.Pop()
	var x int
	if c == '+' {
		x = a + b
	} else if c == '-' {
		x = a - b
	} else if c == '*' {
		x = a * b
	} else if c == '/' {
		x = a / b
	}
	num.Push(x)
}

func isDigit(x rune) bool {
	return x >= '0' && x <= '9'
}
