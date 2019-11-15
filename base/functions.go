package base

import "fmt"

func add(x int, y int) int {
	return x + y
}

func reduce(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(reduce(13, 3))
}
