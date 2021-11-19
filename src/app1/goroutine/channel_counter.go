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

	/*for c := range ch {
		num = c
		fmt.Println("get from ch:", c)
	}*/
	for {
		v, ok := <-ch
		if ok {
			num = v
			fmt.Println("get from ch:", num)
		} else {
			fmt.Println("break")
			break
		}
	}

	fmt.Println(num)
}

func incNum(ch chan int) {
	for i := 0; i < 2; i++ {
		ch <- i
		fmt.Println("insert ok:", i)
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
