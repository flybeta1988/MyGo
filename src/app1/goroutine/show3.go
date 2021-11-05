package main

import (
	"fmt"
	"sync"
)

var wgg sync.WaitGroup

func main() {
	t3()
}

func t3() {
	var ids []int

	wgg.Add(2)
	go work32(&ids)
	go work31(&ids)
	wgg.Wait()

	//time.Sleep(5 * time.Second)
	fmt.Println(ids)
}

func work31(ids *[]int) {
	defer wgg.Done()
	for i := 1; i <= 5; i ++ {
		fmt.Print("work1-", i, " start...")
		*ids = append(*ids, i)
		fmt.Println("work1-", i, "end!")
	}
}

func work32(ids *[]int) {
	defer wgg.Done()
	for i := 6; i <= 10; i ++ {
		fmt.Print("work2-", i, " start...")
		*ids = append(*ids, i)
		fmt.Println("work2-", i, "end!")
	}
}
