package main

import "fmt"

func main() {
	test2()
	fmt.Println("aaa")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println("world")
	fmt.Println("hello")
}

func test2() {
	defer fmt.Println(4)
	defer fmt.Println(5)
	defer fmt.Println(6)
}
