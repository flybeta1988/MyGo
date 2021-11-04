package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	t2()
}

func t2() {
	var ids []int
	var ch chan int
	ch = make(chan int, 15)

	waitGroup.Add(2)
	go work2(ch)
	go work1(ch)


	fmt.Printf("%T\n", ch)
	for i := range ch {
		fmt.Print(i)
		ids = append(ids, i)
		fmt.Println("-x")
	}
	close(ch)
	waitGroup.Wait()
	fmt.Println(ids)
}

func work1(ch chan int ) {
	//time.Sleep(time.Duration(2)*time.Second)
	defer waitGroup.Done()
	for i := 1; i <= 5; i ++ {
		ch <- i
	}
}

func work2(ch chan int) {
	//time.Sleep(time.Duration(2)*time.Second)
	defer waitGroup.Done()
	for i := 6; i <= 10; i ++ {
		ch <- i
	}
}
