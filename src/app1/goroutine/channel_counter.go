package main

import (
	"fmt"
)

var (
	num int
)

func main() {
	var ch = make(chan int)
	//var ch2 = make(chan bool)

	go incNum(ch)

	for c := range ch {
		num = c
		fmt.Println("get from ch:", c)
	}

	fmt.Println(num)
}

func incNum(ch chan int) {
	var n int
	for i := 0; i < 3; i++ {
		n ++
		ch <- n
	}
	close(ch)
}

func read(ch chan int, ch2 chan bool) {
	for {
		v, ok := <- ch
		if ok {
			fmt.Println("read num:", v)
			num = v
		} else {
			ch2 <- true
			fmt.Println("not read num:", v)
		}
	}
}
