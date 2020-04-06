package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\n", e)
		}
	}()
	badCall()
	fmt.Println("After bad call \n")
}

func main() {
	fmt.Println("Starting call test ...")
	test()
	fmt.Println("Ending call test!")
}
