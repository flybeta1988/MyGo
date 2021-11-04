package main

import (
	"fmt"
	"sync"
)

func main() {
	t2()
}

func t2() {
	var ids []int
	var ch chan int
	ch = make(chan int, 15)
	var wg sync.WaitGroup

	wg.Add(3)
	go work2(wg, ch)
	go work1(wg, ch)
	go getIds(wg, ch, &ids)

	fmt.Println("ids:", ids)

	wg.Wait()
	close(ch)
}

func getIds(wg sync.WaitGroup, ch chan int, ids *[]int)  {
	defer wg.Done()
	var i int
	for j := 0; j < 20; j++ {
		select {
		case i = <-ch:
			fmt.Println(i)
			*ids = append(*ids, i)
		default:
			fmt.Println("for default")
			break
		}
		/*if 0 == i {
			fmt.Println("i = 0, break")
			break
		}*/
	}
}

func work1(wg sync.WaitGroup, ch chan int) {
	//time.Sleep(time.Duration(2)*time.Second)
	defer wg.Done()
	for i := 1; i <= 5; i ++ {
		ch <- i
	}
	fmt.Println("work1 ok!")
}

func work2(wg sync.WaitGroup, ch chan int) {
	//time.Sleep(time.Duration(2)*time.Second)
	defer wg.Done()
	for i := 6; i <= 10; i ++ {
		ch <- i
	}
	fmt.Println("work2 ok!")
}
