package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count3 int32
	wg3 sync.WaitGroup
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
		atomic.AddInt32(&count3, 1) //安全的对count2加1
		runtime.Gosched()
	}
}
