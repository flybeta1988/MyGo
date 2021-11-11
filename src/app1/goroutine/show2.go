package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup
func main() {
	t2()
}

func t2() {
	startTime := time.Now().Unix()
	var ids []int
	var ch, ok1, ok2 chan int
	ch = make(chan int, 15)
	ok1 = make(chan int)
	ok2 = make(chan int)

	//waitGroup.Add(3)
	go work2(ch, ok2)
	go work1(ch, ok1)
	go closeChan(ch, ok1, ok2)
	//waitGroup.Wait()
	fmt.Printf("ch cap:%d len:%d\n", cap(ch), len(ch))
	for i := range ch {
		fmt.Print(i)
		ids = append(ids, i)
		fmt.Println("-x")
	}
	fmt.Println(ids)

	endTime := time.Now().Unix()
	fmt.Printf("timeCost: %ds\n", endTime - startTime)
}

func closeChan(ch chan int, ok1 chan int, ok2 chan int) {
	var i, j int
	for {
		select {
		case i = <- ok1:
			fmt.Println("work1 finished!")
		case j = <- ok2:
			fmt.Println("work2 finished!")
		default:
			//fmt.Println("closeChan default!")
		}

		if 2 == i + j {
			close(ch)
			fmt.Println("all work finished, so chan close!")
			return
		}
	}
}

func work1(ch chan int, ok chan int) {
	time.Sleep(time.Duration(3)*time.Second)
	for i := 1; i <= 5; i ++ {
		fmt.Print("work1-", i, " start...")
		ch <- i
		fmt.Println("work1-", i, " end!")
	}

	ok <- 1
	fmt.Println("work1 =====> ok!")
}

func work2(ch chan int, ok chan int) {
	time.Sleep(time.Duration(3)*time.Second)
	for i := 6; i <= 10; i ++ {
		fmt.Print("work2-", i, " start...")
		ch <- i
		fmt.Println("work2-", i, " end!")
	}

	ok <- 1
	fmt.Println("work2 =====> ok!")
}
