package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
}
