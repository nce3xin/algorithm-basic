package main

import "fmt"

const N int = 1e4 + 10

var stk [N]rune
var tt int = 0

func main() {
	var s string
	fmt.Scanf("%s", &s)
	n := len(s)
	a := []rune(s)
	ok := true
	for i := 0; i < n; i++ {
		c := a[i]
		if c == '<' || c == '(' || c == '[' || c == '{' { // 是左括号
			push(c)
		} else { // 是右括号
			if empty() { // 如果栈是空的，来了一个右括号，明显不合法
				ok = false
				break
			} else if !empty() && top() == left(c) {
				pop()
			}
		}
	}
	if !empty() {
		ok = false
	}
	if ok {
		fmt.Printf("yes\n")
	} else {
		fmt.Printf("no\n")
	}
}

// 求右括号所匹配的左括号
func left(x rune) rune {
	if x == '>' {
		return '<'
	} else if x == ')' {
		return '('
	} else if x == ']' {
		return '['
	} else if x == '}' {
		return '{'
	}
	return ' '
}

func push(x rune) {
	tt++
	stk[tt] = x
}

func pop() {
	tt--
}

func top() rune {
	return stk[tt]
}

func empty() bool {
	if tt > 0 {
		return false
	}
	return true
}
