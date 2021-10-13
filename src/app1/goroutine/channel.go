package main

import "fmt"

func Count2(ch chan int) {
	ch <- 1
	fmt.Println("Counting...")
}

func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count2(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
