package main

import (
	"fmt"
	"time"
)

func main() {
	testTimeOut()
}

//@todo 未成功执行
func testTimeOut() {
	ch := make(chan int, 10)
	timeout := make(chan bool, 1)
	ch <- 1

	go func() {
		time.Sleep(1e9) //等待1秒
		timeout <- true
	}()

	select {
		case <-ch:
			fmt.Println("ch pop ok!")
		case <-timeout:
			fmt.Println("timeout pop ok!")
		default:
			fmt.Println("this is default!")
	}
}
