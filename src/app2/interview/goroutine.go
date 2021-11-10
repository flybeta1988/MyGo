package interview

import (
	"fmt"
	"runtime"
	"sync"
)

//考点：go执行的随机性和闭包
func TestWaitGroup() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		fmt.Println("#", i)
		go func() {
			fmt.Println("-i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("+i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
