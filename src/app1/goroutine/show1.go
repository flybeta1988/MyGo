package main

import (
	"fmt"
	"time"
)

func main() {
	tt1()
}

func tt1() {
	var ids []int
	work11(&ids)
	work22(&ids)
	fmt.Println(ids)
}

func work11(ids *[]int) {
	time.Sleep(time.Duration(2)*time.Second)
	for i := 1; i <= 5; i ++ {
		*ids = append(*ids, i)
	}
}

func work22(ids *[]int) {
	time.Sleep(time.Duration(2)*time.Second)
	for i := 6; i <= 10; i ++ {
		*ids = append(*ids, i)
	}
}
