package main

import "fmt"

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	aa = 1 << iota
	bb = 1 << iota
	cc = 1 << iota
)

func main() {
	fmt.Println("c0:", c0)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)

	fmt.Println("a:", aa)
	fmt.Println("b:", bb)
	fmt.Println("c:", cc)
}
