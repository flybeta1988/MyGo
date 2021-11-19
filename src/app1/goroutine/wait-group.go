package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		//time.Sleep(2*time.Second)
		runtime.Gosched() //让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行，目的是为了使资源竞争的结果更明显
		value ++
		count = value
	}
}
