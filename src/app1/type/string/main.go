package main

import (
	"fmt"
)

func main() {
	//testBase()
	testStrFor()
	//testStrFor2()
}

func testStrFor2() {
	str := "ABC,世界"
	for i, ch := range str {
		//以Unicode字符方式遍历时，每个字符的类型是 rune （早期的Go语言用 int 类型表示Unicode字符），而不是byte
		//fmt.Println(i, ch)
		fmt.Printf("%d, %c\n", i, ch)
	}
}

func testStrFor() {
	str := "ABC,世界"
	num := len(str)
	for i := 0; i < num; i ++ {
		ch := str[i] //依据下标取字符串中的字符，类型为byte
		fmt.Printf("%c\n", ch)
		//fmt.Println(i, ch)
	}
}

func testBase() {
	hello := "hello"
	world := "world"
	str := "你好，" + hello + world
	fmt.Printf("str len:%d\n", len(str))
	fmt.Printf("str first char:%c\n", str[0])
}
