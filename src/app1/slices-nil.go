package main

import "fmt"

//slice 的零值是 `nil`。
//一个 nil 的 slice 的长度和容量是 0。
func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))

	if nil == z {
		fmt.Println("nil")
	}
}
