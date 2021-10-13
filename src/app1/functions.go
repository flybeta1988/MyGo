package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func reduce(x, y int) int {
	return x - y
}

func foo(args ...int)  {
	bar(args...)
}

func bar(args ...int) {
	for _, arg := range args{
		fmt.Println(arg)
	}
	fmt.Println("++++")
}

func main() {
	/*fmt.Println(add(2, 3))
	fmt.Println(reduce(13, 3))*/
	//foo(1, 2, 3)
	foo(1, 2, 3, 4)
	//bar(1, 2, 3, 4, 5, "aaa")
}
