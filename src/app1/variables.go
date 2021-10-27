package main

import "fmt"

func main() {
	test1()
}

//意外的变量幽灵
func test1() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}
