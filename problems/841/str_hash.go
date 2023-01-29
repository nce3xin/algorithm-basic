package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N int = 1e5 + 10

type ULL uint64 // unsigned long long

var h, p [N]ULL // h[k]存储字符串前k个字母的哈希值, p[k]存储 P^k mod 2^64
var P ULL = 131

func main() {
	r := bufio.NewReader(os.Stdin)
	in := readInt(r)
	n, m := in[0], in[1]
	s, _ := r.ReadString('\n')

	s = " " + s // 因为char是从下标为1的地方开始使用的，所以这里偏移1位
	char := []rune(s)
	p[0] = 1 // p的0次方等于1
	for i := 1; i <= n; i++ {
		// 因为要用到p的次方，所以用p[]把p的所有次方都预处理出来
		p[i] = p[i-1] * P
		h[i] = h[i-1]*P + ULL(char[i]) // char[i]只要不是0就可以，是多少都无所谓
	}

	for ; m > 0; m-- {
		in = readInt(r)
		l1, r1, l2, r2 := in[0], in[1], in[2], in[3]
		if get(l1, r1) == get(l2, r2) {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}
}

func get(l, r int) ULL {
	return h[r] - h[l-1]*p[r-l+1]
}

func readInt(r *bufio.Reader) []int {
	s, _ := r.ReadString('\n')
	ss := strings.Fields(s)
	res := make([]int, len(ss))
	for i, v := range ss {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}
