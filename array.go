package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(len(a))
}
