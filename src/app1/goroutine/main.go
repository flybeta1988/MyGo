package main

import (
	"fmt"
	"sync"
)

var wg5 sync.WaitGroup

func Add(x, y int)  {
	defer wg5.Done()
	z := x + y
	fmt.Println(z)
}

func main() {
	wg5.Add(10)
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	wg5.Wait()
}
