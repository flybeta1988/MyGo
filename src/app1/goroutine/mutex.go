package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count2 int32
	wg2 sync.WaitGroup
)

func main() {
	wg2.Add(2)
	go incCount2()
	go incCount2()
	wg2.Wait()
	fmt.Println(count2)
}

func incCount2() {
	defer wg2.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt32(&count2, 1) //安全的对count2加1
		runtime.Gosched()
	}
}
