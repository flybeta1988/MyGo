package main

import "fmt"

func main() {
	//testBase()
	testOther()
}

func testBase() {
	var v1 bool
	v1 = true
	v2 := (1 == 2)
	fmt.Printf("type:%T, value:%v\n", v1)
	fmt.Printf("type:%T, value:%v\n", v2)
}

//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
func testOther() {
	var b bool
	//b = 1 //编译错误
	//b = bool(1) //编译错误
	b = (1 != 0)
	fmt.Println("Result:", b)
}
