package main

import (
	"fmt"
)

type OneWayChannel struct {}

func main() {
	owc := &OneWayChannel{}
	owc.test1()
}

func (owc OneWayChannel) test1() {

	var ch = make(chan int)
	var ids = make(chan int)

	go owc.write(ch)
	go owc.read(ch, ids)

	for j := range ids {
		fmt.Println("j:", j)
	}

	fmt.Println("finished!")
}

func (OneWayChannel) write(wch chan<- int) {
	for i := 0; i < 5; i ++ {
		wch <- i
	}
	close(wch)
}

func (OneWayChannel) read(rch <-chan int, ids chan<- int) {
	for i := range rch {
		ids <- i
	}
	close(ids)
}
