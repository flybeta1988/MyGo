package main

import "fmt"

func main() {
	test1()
}

func test1() {
	var b byte
	b = 1
	str := "abcABC"
	fmt.Printf("%d, %T\n", b, b)
	for k, v := range str {
		fmt.Printf("%d, %T, %d\n", k, v, v)
		//fmt.Println(k, v)
	}
}

func test2[T int | string](a, b T) T {
	return a + b
}
