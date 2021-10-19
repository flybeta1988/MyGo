package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count3 int32
	wg3 sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	fmt.Println("CPU nums:", runtime.NumCPU())
	wg3.Add(2)
	go incCount3()
	go incCount3()
	wg3.Wait()
	fmt.Println(count3)
}

func incCount3() {
	defer wg3.Done()
	for i := 0; i < 2; i++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := count3
			runtime.Gosched()
			value ++
			count3 = value
		}
		mutex.Unlock()
	}
}
