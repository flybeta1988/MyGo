package main

import (
	"fmt"
)

func main() {
	//test2()
	//testSelect()
	testChanFor()
}

func testChanFor() {
	ch := make(chan int, 5)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	fmt.Println(len(ch))
	fmt.Printf("ch type:%T\n", ch)

	for i := 0; i < 10; i ++ {
		select {
		case j, ok := <-ch:
			fmt.Println(j, " ", ok)
		default:
			fmt.Println("this is default!")
		}
	}
}

func testSelect() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	select {
		case a := <-ch1:
			fmt.Println("ch1 pop ok!", a)
		case ch2 <- 2:
			fmt.Println("ch2 push ok!")
		default:
			fmt.Println("this is default!")
	}
	close(ch1)
	close(ch2)
}

func test2()  {
	//var intChan chan int
	intChan := make(chan int, 2)
	intChan <- 1
	intChan <- 2
	id := <-intChan
	id2 := <-intChan
	fmt.Println(id)
	fmt.Println(id2)
}

func Count2(ch chan int) {
	ch <- 1
	ch <- 2
	fmt.Println("Counting...")
}

func test1() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count2(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
