package main

import (
	"fmt"
	"sync"
)

var wg4 sync.WaitGroup

func main() {
	//test1()
	//test2()
	test3()
	//test4()
	//testSelect()
	//testChanFor()
}

func testChanFor() {
	ch := make(chan int, 5)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	//fmt.Println(len(ch))
	fmt.Printf("ch type:%T\n", ch)

	size := len(ch)
	for i := 0; i < size; i ++ {
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
	ch3 := make(chan int, 2)
	ch1 <- 1

	go mySleep(ch3)

	for {
		v, ok := <-ch3
		if ok {
			fmt.Println("ch3 pop v:", v)
		} else {
			fmt.Println("for break")
			break
		}
	}
}

func mySleep(ch chan int) {
	//defer wg4.Done()
	//time.Sleep(1 * time.Second)
	ch <- 3
	close(ch)
}

func test2()  {
	//var intChan chan int
	intChan := make(chan int, 2)
	intChan <- 1
	intChan <- 2
	id := <-intChan
	id2 := <-intChan
	fmt.Println(id, id2)
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

func test3() {
	var ch = make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	/*for c := range ch {
		fmt.Println("for c:", c)
	}*/

	for  {
		if v, ok := <-ch; ok {
			fmt.Println("v:", v)
		} else {
			fmt.Println("not ok")
			break;
		}
	}
}

func test4() {
	var ids []int
	var ch = make(chan int)

	var wg6 sync.WaitGroup
	wg6.Add(2)

	go func() {
		defer wg6.Done()
		size := len(ch)
		for j := 0; j < size; j ++ {
			select {
			case t := <-ch:
				fmt.Println("read t:", t)
				ids = append(ids, t)
			default:
				fmt.Println("select default")
			}
		}
	}()

	for i := 0; i < 5; i ++ {
		//ch <- i
		fmt.Println("add ok i:", i)
	}
	//close(ch)
	wg6.Wait()
}
