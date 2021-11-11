package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count2 int32
	wg2 sync.WaitGroup
	mutex sync.Mutex
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
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := count2
			runtime.Gosched()
			value ++
			count2 = value
		}
		mutex.Unlock()
	}
}
