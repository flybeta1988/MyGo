package main

import "fmt"

func main() {
	test1()
}

func test1() {
	var b byte
	b = 1
	str := "abc"
	fmt.Printf("%d, %T\n", b, b)
	for k, v := range str {
		fmt.Printf("%d, %T\n", k, v)
		//fmt.Println(k, v)
	}
}
