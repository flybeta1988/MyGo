package main

import "fmt"

func main() {
	var str string
	var foo interface{}

	//foo = 100
	fmt.Println(foo)

	foo = "hello null interface"
	fmt.Println(foo)

	str, ok := foo.(string) //需要类型断言
	fmt.Printf(str)
	fmt.Printf("type assignment status: %t", ok)
}
