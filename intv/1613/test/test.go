package main

import "fmt"

func main() {
	g := [2][2]byte{{'1', '2'}, {'3', '4'}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%v ", g[i][j]-'0')
		}
		fmt.Println()
	}

	fmt.Println()

	a := 22
	g[0][0] = byte(a) + '0'
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%v ", g[i][j]-'0')
		}
		fmt.Println()
	}
}
