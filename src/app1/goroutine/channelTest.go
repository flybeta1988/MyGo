package main

import "fmt"

func main() {
	test11()
}

func test11() {
	//var chInt chan int
	chInt := make(chan int, 1)
	chInt <- 1
	//go add(chInt)
	i := <- chInt
	fmt.Println(i)
}

func add(ch chan int) {
	ch <- 1
}
